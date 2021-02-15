package godd

import (
	"log"
	"time"
)

// Portal Struct
type Portal struct {
	appList    []appServe
	iCloseList []InterfaceClose
	// router fiber.Router
}

type appServe struct {
	app  InterfaceApp
	port string
}

// NewPortal Func
func NewPortal() *Portal {
	return &Portal{}
}

// AppendApp Func
func (pt *Portal) AppendApp(app interface{}, port string, framework ...FrameWork) {

	interfaceApp := AdapterApp(app, framework...)
	addAPIGetHealth(interfaceApp)

	pt.appList = append(pt.appList, appServe{
		app:  interfaceApp,
		port: port,
	})
}

func startAppGoroutine(app appServe, errc chan error) {
	go func() {
		var ch chan error
		log.Println("App Start")
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
	log.Println("App Shutdown")
	MustError(app.app.Shutdown())
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
	defer pt.deferInterfaceClose()

	// Running Server Wait for Error
	waitAppAppGoroutine(errc, 1)

	for _, app := range pt.appList {
		shutdownAppGoroutine(app)
	}
}

//================================================

// AppendInterfaceClose Func
func (pt *Portal) AppendInterfaceClose(iList ...interface{}) {
	for _, i := range iList {
		pt.iCloseList = append(pt.iCloseList, i.(InterfaceClose))
	}
}

func (pt *Portal) deferInterfaceClose() {
	for _, iClose := range pt.iCloseList {
		iClose.Close()
	}
}
