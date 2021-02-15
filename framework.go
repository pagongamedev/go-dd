package godd

// AdapterApp Func
func AdapterApp(app interface{}, framework ...FrameWork) InterfaceApp {

	fw := FrameWorkGofiber
	if len(framework) > 0 {
		fw = framework[0]
	}
	var interfaceApp InterfaceApp

	switch fw {
	case FrameWorkGofiber:
		interfaceApp = AdapterAppGofiber(app, fw)
		break
	}

	return interfaceApp
}
