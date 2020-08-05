package appkit

import (
	"goco"
	"goco/foundation"
	"unsafe"
)

type NSWindow struct {
	goco.NSObject
}

func (win NSWindow) Alloc() NSWindow {
	win.NSObject = goco.NSObject{Ptr: goco.InvokeAndReturn("NSWindow", "alloc", goco.NSClassFromString("NSWindow"))}
	return win
}

func (win NSWindow) InitWithContentRectStyleMaskBackingDefer(contentRect foundation.NSRect, styleMask NSWindowStyleMask, storeType NSBackingStoreType, doDefer bool) NSWindow {
	selector := "initWithContentRect:styleMask:backing:defer:"
	inv := goco.InvocationWithMethodSignature("NSWindow", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&contentRect), 2)
	inv.SetArgumentAtIndex(unsafe.Pointer(&styleMask), 3)
	inv.SetArgumentAtIndex(unsafe.Pointer(&storeType), 4)
	inv.SetArgumentAtIndex(unsafe.Pointer(&doDefer), 5)
	inv.InvokeWithTarget(win.Ptr)
	return win
}

func (win NSWindow) CascadeTopLeftFromPoint(topLeftPoint foundation.NSPoint) foundation.NSPoint {
	selector := "cascadeTopLeftFromPoint:"
	inv := goco.InvocationWithMethodSignature("NSWindow", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&topLeftPoint), 2)
	inv.InvokeWithTarget(win.Ptr)
	return foundation.NSPoint{} //TODO: get return NSPoint
}

func (win NSWindow) SetTitle(title foundation.NSString) {
	goco.Invoke("NSWindow", "setTitle:", win.Ptr, unsafe.Pointer(&title.Ptr))
}

func (win NSWindow) MakeKeyAndOrderFront() { //TODO: take in sender id
	selector := "makeKeyAndOrderFront:"
	inv := goco.InvocationWithMethodSignature("NSWindow", selector)
	inv.SetSelector(selector)
	obj := goco.NSObject{}
	inv.SetArgumentAtIndex(unsafe.Pointer(&obj), 2)
	inv.InvokeWithTarget(win.Ptr)
}
