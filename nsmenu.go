package goco

import "unsafe"

type NSMenu struct {
	NSObject
}

func (menu NSMenu) Alloc() NSMenu {
	selector := "alloc"
	inv := InvocationWithMethodSignature("NSMenu", selector)
	inv.SetSelector(selector)
	inv.InvokeWithTarget(NSClassFromString("NSMenu"))
	menu.NSObject = NSObject{inv.GetReturnValue()}
	return menu
}

func (menu NSMenu) AddItem(item NSMenuItem) {
	selector := "addItem:"
	inv := InvocationWithMethodSignature("NSMenu", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&item), 2)
	inv.InvokeWithTarget(menu.ptr)
}
