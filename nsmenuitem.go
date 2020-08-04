package goco

import "unsafe"

type NSMenuItem struct {
	NSObject
}

func (item NSMenuItem) Alloc() NSMenuItem {
	item.NSObject = NSObject{InvokeAndReturn("NSMenuItem", "alloc", NSClassFromString("NSMenuItem"))}
	return item
}

func (item NSMenuItem) InitWithTitleActionKeyEquivalent(title NSString, sel string, keyEq NSString) NSMenuItem {
	selector := "initWithTitle:action:keyEquivalent:"
	inv := InvocationWithMethodSignature("NSMenuItem", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&title.ptr), 2)
	selVar := NSSelectorFromString(sel)
	inv.SetArgumentAtIndex(unsafe.Pointer(&selVar), 3)
	inv.SetArgumentAtIndex(unsafe.Pointer(&keyEq.ptr), 4)
	inv.InvokeWithTarget(item.ptr)
	return item
}

func (item NSMenuItem) SetSubMenu(menu NSMenu) {
	selector := "setSubmenu:"
	inv := InvocationWithMethodSignature("NSMenuItem", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&menu.ptr), 2)
	inv.InvokeWithTarget(item.ptr)
}
