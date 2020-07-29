package appkit

import (
	"goco"
	"goco/foundation"
	"unsafe"
)

type NSWindow struct {
	foundation.NSObject
}

func (win NSWindow) Alloc() NSWindow {
	win.NSObject = foundation.PerformSelectorClass("NSWindow", "alloc")
	return win
}

func (win NSWindow) InitWithContentRectStyleMaskBackingDefer(contentRect foundation.NSRect, styleMask NSWindowStyleMask, storeType NSBackingStoreType, doDefer bool) NSWindow {
	/*void* NSWindow_initWithContentRectstyleMaskbackingdefer(int x, int y, int width, int height, int mask, int storetype, unsigned char defer) {
		return [[NSWindow alloc] initWithContentRect:NSMakeRect(x, y, width, height)
	        styleMask:mask backing:storetype defer:defer];
	}
	func InitWindow(x, y, width, height int, styleMask NSWindowStyleMask, storeType NSBackingStoreType, doDefer bool) NSWindow {
		return NSWindow{NSObject: foundation.NSObject{Ptr: C.NSWindow_initWithContentRectstyleMaskbackingdefer(C.int(x), C.int(y), C.int(width), C.int(height), C.int(styleMask), C.int(storeType), C.uchar(goco.BoolToBOOL(doDefer)))}}
	}
	*/
	selector := "initWithContentRect:styleMask:backing:defer:"
	sig := win.MethodSignatureForSelector(selector)
	inv := foundation.InvocationWithMethodSignature(sig)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&contentRect), 2)
	inv.SetArgumentAtIndex(unsafe.Pointer(&styleMask), 3)
	inv.SetArgumentAtIndex(unsafe.Pointer(&storeType), 4)
	d := goco.BoolToBOOL(doDefer)
	inv.SetArgumentAtIndex(unsafe.Pointer(&d), 5)
	inv.InvokeWithTarget(win.NSObject)
	return win
}

func (win NSWindow) CascadeTopLeftFromPoint(topLeftPoint foundation.NSPoint) foundation.NSPoint {
	//Relies heavily on this: https://stackoverflow.com/questions/904515/how-to-use-performselectorwithobjectafterdelay-with-primitive-types-in-cocoa
	selector := "cascadeTopLeftFromPoint:" //foundation.NSSelectorFromString("cascadeTopLeftFromPoint:")
	sig := win.MethodSignatureForSelector(selector)
	inv := foundation.InvocationWithMethodSignature(sig)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&topLeftPoint.Point), 2)
	inv.InvokeWithTarget(win.NSObject)
	//return foundation.NSPoint{Point: *(*C.NSPoint)(inv.GetReturnValue())}
	return foundation.NSPoint{} //TODO: get return NSPoint
}

func (win NSWindow) SetTitle(title foundation.NSString) {
	win.PerformSelectorWithObject("setTitle:", title.NSObject)
}

func (win NSWindow) MakeKeyAndOrderFront() { //TODO: take in sender id
	win.PerformSelectorWithObject("makeKeyAndOrderFront:", foundation.NSObject{})
}
