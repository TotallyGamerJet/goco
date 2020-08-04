package goco

import (
	"unsafe"
)

type NSWindow struct {
	NSObject
}

func (win NSWindow) Alloc() NSWindow {
	selector := "alloc"
	inv := InvocationWithMethodSignature("NSWindow", selector)
	inv.SetSelector(selector)
	inv.InvokeWithTarget(NSClassFromString("NSWindow"))
	win.NSObject = NSObject{inv.GetReturnValue()}
	return win
}

func (win NSWindow) InitWithContentRectStyleMaskBackingDefer(contentRect NSRect, styleMask NSWindowStyleMask, storeType NSBackingStoreType, doDefer bool) NSWindow {
	selector := "initWithContentRect:styleMask:backing:defer:"
	inv := InvocationWithMethodSignature("NSWindow", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&contentRect), 2)
	inv.SetArgumentAtIndex(unsafe.Pointer(&styleMask), 3)
	inv.SetArgumentAtIndex(unsafe.Pointer(&storeType), 4)
	d := BoolToBOOL(doDefer)
	inv.SetArgumentAtIndex(unsafe.Pointer(&d), 5)
	inv.InvokeWithTarget(win.ptr)
	return win
}

func (win NSWindow) CascadeTopLeftFromPoint(topLeftPoint NSPoint) NSPoint {
	//Relies heavily on this: https://stackoverflow.com/questions/904515/how-to-use-performselectorwithobjectafterdelay-with-primitive-types-in-cocoa
	selector := "cascadeTopLeftFromPoint:"
	inv := InvocationWithMethodSignature("NSWindow", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&topLeftPoint), 2)
	inv.InvokeWithTarget(win.ptr)
	return NSPoint{} //TODO: get return NSPoint
}

func (win NSWindow) SetTitle(title NSString) {
	selector := "setTitle:"
	inv := InvocationWithMethodSignature("NSWindow", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&title.ptr), 2)
	inv.InvokeWithTarget(win.ptr)
}

func (win NSWindow) MakeKeyAndOrderFront() { //TODO: take in sender id
	selector := "makeKeyAndOrderFront:"
	inv := InvocationWithMethodSignature("NSWindow", selector)
	inv.SetSelector(selector)
	obj := NSObject{}
	inv.SetArgumentAtIndex(unsafe.Pointer(&obj), 2)
	inv.InvokeWithTarget(win.ptr)
}
