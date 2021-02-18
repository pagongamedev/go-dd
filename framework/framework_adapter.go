package framework

import godd "github.com/pagongamedev/go-dd"

// AdapterApp Func
func AdapterApp(app interface{}, frameworkType ...godd.FrameWork) godd.InterfaceApp {

	fwType := godd.FrameWorkGofiber
	if len(frameworkType) > 0 {
		fwType = frameworkType[0]
	}
	var interfaceApp godd.InterfaceApp

	switch fwType {
	case godd.FrameWorkGofiber:
		interfaceApp = AdapterAppGofiber(app, fwType)
		break
	}

	return interfaceApp
}
