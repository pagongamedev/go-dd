package portal

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	godd "github.com/pagongamedev/go-dd"
	"github.com/pagongamedev/go-dd/framework"
)

// Portal Struct
type Portal struct {
	appList        []appServe
	deferCloseList []godd.DeferClose
	// router fiber.Router
}

// New Func
func New(secret *godd.MapString, funcEnvironment godd.FuncEnvironment, funcMiddleware godd.FuncEnvironment) (*Portal, *godd.MapString, *godd.Map, *godd.Map) {
	portal := Portal{}
	var stateEnvironment *godd.Map
	var stateMiddleware *godd.Map
	var deferCloseList *[]godd.DeferClose
	if secret != nil {
		if funcEnvironment != nil {
			stateEnvironment, deferCloseList = funcEnvironment(*secret)
			portal.AppendDeferClose((*deferCloseList)...)
		}
		if funcMiddleware != nil {
			stateMiddleware, deferCloseList = funcMiddleware(*secret)
			portal.AppendDeferClose((*deferCloseList)...)
		}
	}

	return &portal, secret, stateEnvironment, stateMiddleware
}

type appServe struct {
	app  godd.InterfaceApp
	port string
}

// AppendApp Func
func (pt *Portal) AppendApp(app interface{}, port string, fw ...godd.FrameWork) {

	interfaceApp := framework.AdapterApp(app, fw...)
	addAPIGetHealth(interfaceApp)

	pt.appList = append(pt.appList, appServe{
		app:  interfaceApp,
		port: port,
	})
}

func startAppGoroutine(app appServe, errc chan error) {
	go func() {
		var ch chan error
		err := app.app.Listen(app.port)
		if err != nil {
			ch = errc
		}
		select {
		case ch <- err:
			return
		default:

		}

	}()
}

func shutdownAppGoroutine(app appServe) {
	err := app.app.Shutdown()
	godd.MustError(err)
	log.Println("Shutdown App " + app.port)
}

func waitAppGoroutine(errc chan error, done chan bool, waitTimeForError int64) {
	var timeEnd int64
	for timeEnd == 0 || timeEnd > time.Now().Unix() {

		select {
		case err := <-errc:
			log.Println(err)
			if timeEnd == 0 {
				timeEnd = time.Now().Unix() + waitTimeForError
			}
		case done := <-done:
			if done == true {
				timeEnd = time.Now().Unix() + waitTimeForError
				break
			}
		default:

		}
	}
}

// StartServer Func
func (pt *Portal) StartServer() {
	errc := make(chan error)
	done := make(chan bool, 1)
	for _, app := range pt.appList {
		startAppGoroutine(app, errc)
	}

	// Defer InterfaceClose
	defer pt.deferQuitApp()

	// Check Defer when Ctrl+C in Command
	pt.checkInterruptQuit(done)

	// Running Server Wait for Error
	waitAppGoroutine(errc, done, 1)

	return
}

//================================================

// AppendDeferClose Func
func (pt *Portal) AppendDeferClose(iList ...godd.DeferClose) {
	pt.deferCloseList = append(pt.deferCloseList, iList...)
}

func (pt *Portal) deferClose() {
	log.Printf("Start Defer Close\n")
	for _, d := range pt.deferCloseList {
		err := d.I.Close()
		if err != nil {
			log.Printf("InterfaceClose %v Error : %v\n", d.Name, err)
		}
		log.Printf("Defer Close : %v\n", d.Name)
	}
}

func (pt *Portal) checkInterruptQuit(done chan bool) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		for sig := range c {
			if sig != nil {
				log.Printf("\nClose With Ctrl C\n")
				pt.deferQuitApp()
				done <- true
				// os.Exit(1)
			}
		}
	}()
}

//=====================================
func (pt *Portal) deferQuitApp() {
	for _, app := range pt.appList {
		shutdownAppGoroutine(app)
	}

	pt.deferClose()
	log.Println("End Portal")
}

//=====================================

// addAPIGetHealth Func
func addAPIGetHealth(app godd.InterfaceApp) {
	app.Get("/health", handlerHealth())
}

func handlerHealth() godd.Handler {
	return func(ctx godd.InterfaceContext) error {
		return ctx.Response(godd.Map{"success": true})
	}
}
