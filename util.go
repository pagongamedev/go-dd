package godd

import (
	"log"
)

// =====================================================================
//                              Add On
// =====================================================================

// MustError Func
func MustError(err error, strList ...string) {
	if err != nil {
		if strList != nil {
			log.Fatal(strList)
		} else {
			log.Fatal("Error : ", err)
		}
	}
}

// AddAPIGetHealth Func
func addAPIGetHealth(app InterfaceApp) {
	app.Get("/health", handlerHealth())
}

func handlerHealth() Handler {
	return func(ctx InterfaceContext) error {
		return ctx.Response(Map{"success": true})
	}
}
