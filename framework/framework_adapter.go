package framework

import godd "github.com/pagongamedev/go-dd"

// AdapterApp Func
func AdapterApp(app interface{}, frameworkType ...godd.FrameWork) godd.InterfaceApp {

	fwType := godd.FrameWorkGofiberV2
	if len(frameworkType) > 0 {
		fwType = frameworkType[0]
	}
	var interfaceApp godd.InterfaceApp

	switch fwType {
	case godd.FrameWorkGofiberV2:
		interfaceApp = AdapterAppGofiber(app, fwType)
		break
	}

	return interfaceApp
}
