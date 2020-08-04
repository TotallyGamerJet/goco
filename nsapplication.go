package goco

import (
	"unsafe"
)

type NSApplication struct {
	NSObject
}

func (app NSApplication) SharedApplication() NSApplication {
	app.NSObject = NSObject{InvokeAndReturn("NSApplication", "sharedApplication", NSClassFromString("NSApplication"))}
	return app
}

func (app NSApplication) SetActivationPolicy(policy NSApplicationActivationPolicy) bool {
	selector := "setActivationPolicy:"
	inv := InvocationWithMethodSignature("NSApplication", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&policy), 2)
	inv.InvokeWithTarget(app.ptr)
	return *(*bool)(inv.GetReturnValue())
}

func (app NSApplication) ActivateIgnoringOtherApps(flag bool) {
	selector := "activateIgnoringOtherApps:"
	inv := InvocationWithMethodSignature("NSApplication", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&flag), 2)
	inv.InvokeWithTarget(app.ptr)
}

func (app NSApplication) SetMainMenu(menu NSMenu) {
	selector := "setMainMenu:"
	inv := InvocationWithMethodSignature("NSApplication", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&menu.ptr), 2)
	inv.InvokeWithTarget(app.ptr)
}

func (app NSApplication) Run() {
	Invoke("NSApplication", "run", app.ptr)
}

/*
func (app NSApplication) Terminate() { //TODO: take in sender
	app.PerformSelectorWithObject("terminate:", foundation.NSObject{})
}*/
