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

func (app NSApplication) SetAppDelegate(delegate NSApplicationDelegate) {
	panic("implement")
}

func (app NSApplication) SetActivationPolicy(policy NSApplicationActivationPolicy) bool {
	selector := "setActivationPolicy:"
	inv := goco.InvocationWithMethodSignature("NSApplication", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&policy), 2)
	inv.InvokeWithTarget(app.Ptr)
	return *(*bool)(inv.GetReturnValue())
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
