package appkit

import (
	"goco/foundation"
	"unsafe"
)

type NSApplication struct {
	foundation.NSObject
}

func (app NSApplication) SharedApplication() NSApplication {
	app.NSObject = foundation.PerformSelectorClass("NSApplication", "sharedApplication")
	return app
}

func (app NSApplication) SetActivationPolicy(policy NSApplicationActivationPolicy) bool {
	selector := "setActivationPolicy:"
	sig := app.MethodSignatureForSelector(selector)
	inv := foundation.InvocationWithMethodSignature(sig)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&policy), 2)
	inv.InvokeWithTarget(app.NSObject)
	return *(*bool)(inv.GetReturnValue())
}

func (app NSApplication) ActivateIgnoringOtherApps(flag bool) {
	selector := "activateIgnoringOtherApps:"
	sig := app.MethodSignatureForSelector(selector)
	inv := foundation.InvocationWithMethodSignature(sig)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&flag), 2)
	inv.InvokeWithTarget(app.NSObject)
}

func (app NSApplication) SetMainMenu(menu NSMenu) {
	app.PerformSelectorWithObject("setMainMenu:", menu.NSObject)
}

func (app NSApplication) Run() {
	app.PerformSelector("run")
}

func (app NSApplication) Terminate() { //TODO: take in sender
	app.PerformSelectorWithObject("terminate:", foundation.NSObject{})
}
