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
	appList    []appServe
	iCloseList []godd.InterfaceClose
	// router fiber.Router
}

// New Func
func New() *Portal {
	return &Portal{}
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
		}
	}()
}

func shutdownAppGoroutine(app appServe) {
	err := app.app.Shutdown()
	godd.MustError(err)
	log.Println("Shutdown App " + app.port)
}

func waitAppAppGoroutine(errc chan error, waitTimeForError int64) {
	var timeEnd int64
	for timeEnd == 0 || timeEnd > time.Now().Unix() {

		select {
		case err := <-errc:
			log.Println(err)
			if timeEnd == 0 {
				timeEnd = time.Now().Unix() + waitTimeForError
			}
			break
		default:
			break
		}
	}
}

// StartServer Func
func (pt *Portal) StartServer() {
	errc := make(chan error)
	for _, app := range pt.appList {
		startAppGoroutine(app, errc)
	}

	// Defer InterfaceClose
	defer pt.deferQuitApp()

	// Check Defer when Ctrl+C in Command
	pt.checkInterruptQuit()

	// Running Server Wait for Error
	waitAppAppGoroutine(errc, 1)

}

//================================================

// AppendInterfaceClose Func
func (pt *Portal) AppendInterfaceClose(iList ...interface{}) {
	for _, i := range iList {
		pt.iCloseList = append(pt.iCloseList, i.(godd.InterfaceClose))
	}
}

func (pt *Portal) deferInterfaceClose() {
	for _, iClose := range pt.iCloseList {
		err := iClose.Close()
		if err != nil {
			log.Println("InterfaceClose Error:", err)
		}
	}
}

func (pt *Portal) checkInterruptQuit() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		pt.deferQuitApp()
		os.Exit(1)
	}()
}

//=====================================
func (pt *Portal) deferQuitApp() {
	for _, app := range pt.appList {
		shutdownAppGoroutine(app)
	}

	pt.deferInterfaceClose()
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
