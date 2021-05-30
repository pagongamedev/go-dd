package portal

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	godd "github.com/pagongamedev/go-dd"
)

// Portal Struct
type Portal struct {
	appList        []appServe
	deferCloseList []godd.DeferClose
	// router fiber.Router
}

// New Func
func New(secret *godd.MapString, funcStorage godd.FuncEnvironment, funcMiddleware godd.FuncEnvironment) (*Portal, *godd.MapString, *godd.Map, *godd.Map) {
	portal := Portal{}
	var stateStorage *godd.Map
	var stateMiddleware *godd.Map
	var deferCloseList *[]godd.DeferClose
	if secret != nil {
		if funcStorage != nil {
			stateStorage, deferCloseList = funcStorage(*secret)
			portal.AppendDeferClose((*deferCloseList)...)
		}
		if funcMiddleware != nil {
			stateMiddleware, deferCloseList = funcMiddleware(*secret)
			portal.AppendDeferClose((*deferCloseList)...)
		}
	}

	return &portal, secret, stateStorage, stateMiddleware
}

type appServe struct {
	app       godd.InterfaceApp
	port      string
	extraList []interface{}
}

// AppendApp Func
func (pt *Portal) AppendApp(interfaceApp godd.InterfaceApp, port string, extraList ...interface{}) {

	addAPIGetHealth(interfaceApp)

	pt.appList = append(pt.appList, appServe{
		app:       interfaceApp,
		port:      port,
		extraList: extraList,
	})
}

func startAppGoroutine(app appServe, errc chan error) {
	go func() {
		err := app.app.Listen(app.port, app.extraList...)
		if err != nil {
			errc <- err
		}
		select {
		case <-errc:
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
				log.Printf("Close With Ctrl C\n")
				// pt.deferQuitApp()
				// os.Exit(1)
				done <- true
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
	app.Get("/health", nil, handlerHealth())
}

func handlerHealth() godd.Handler {
	return func(context *godd.Context) error {
		return context.App().Response(godd.Map{"success": true}, "")
	}
}
