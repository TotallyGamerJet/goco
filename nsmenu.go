package goco

import "unsafe"

type NSMenu struct {
	NSObject
}

func (menu NSMenu) Alloc() NSMenu {
	menu.NSObject = NSObject{InvokeAndReturn("NSMenu", "alloc", NSClassFromString("NSMenu"))}
	return menu
}

func (menu NSMenu) AddItem(item NSMenuItem) {
	selector := "addItem:"
	inv := InvocationWithMethodSignature("NSMenu", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&item), 2)
	inv.InvokeWithTarget(menu.ptr)
}
