package godd

import (
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"
)

// Portal Struct
type Portal struct {
	appList    []appServe
	eg         errgroup.Group
	iCloseList []InterfaceClose
	// router fiber.Router
}

type appServe struct {
	app  *fiber.App
	port string
}

// NewPortal Func
func NewPortal() *Portal {
	return &Portal{}
}

// AppendApp Func
func (pt *Portal) AppendApp(app *fiber.App, port string) {

	addAPIGetHealth(app)

	pt.appList = append(pt.appList, appServe{
		app:  app,
		port: port,
	})
}

func startAppGoroutine(eg *errgroup.Group, app appServe) {
	eg.Go(func() error {
		return app.app.Listen(app.port)
	})
}

func shutdownAppGoroutine(app appServe) {
	MustError(app.app.Shutdown())
}

// StartServer Func
func (pt *Portal) StartServer() {
	for _, app := range pt.appList {
		startAppGoroutine(&pt.eg, app)
	}

	// Defer InterfaceClose
	defer pt.deferInterfaceClose()

	// Running Server
	MustError(pt.eg.Wait())

	// Closing Server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// for Echo use Context
	// _, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	for _, app := range pt.appList {
		shutdownAppGoroutine(app)
	}
}

//==================================

// InterfaceClose for Manage Defer Close
type InterfaceClose interface {
	Close() error
}

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
