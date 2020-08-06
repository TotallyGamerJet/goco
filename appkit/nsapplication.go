package appkit

import (
	"goco"
	"unsafe"
)

type NSApplication struct {
	goco.NSObject
}

func (app NSApplication) SharedApplication() NSApplication {
	app.NSObject = goco.NSObject{Ptr: goco.InvokeAndReturn("NSApplication", "sharedApplication", goco.NSClassFromString("NSApplication"))}
	return app
}

func (app NSApplication) SetActivationPolicy(policy NSApplicationActivationPolicy) bool {
	return *(*bool)(goco.InvokeAndReturn("NSApplication", "setActivationPolicy:", app.Ptr, unsafe.Pointer(&policy)))
}

func (app NSApplication) ActivateIgnoringOtherApps(flag bool) {
	goco.Invoke("NSApplication", "activateIgnoringOtherApps:", app.Ptr, unsafe.Pointer(&flag))
}

func (app NSApplication) SetMainMenu(menu NSMenu) {
	goco.Invoke("NSApplication", "setMainMenu:", app.Ptr, unsafe.Pointer(&menu.Ptr))
}

func (app NSApplication) Run() {
	goco.Invoke("NSApplication", "run", app.Ptr)
}

/*
func (app NSApplication) Terminate() { //TODO: take in sender
	app.PerformSelectorWithObject("terminate:", foundation.NSObject{})
}*/
