package appkit

import (
	"goco"
	"goco/foundation"
	"unsafe"
)

type NSMenuItem struct {
	goco.NSObject
}

func (item NSMenuItem) Alloc() NSMenuItem {
	item.NSObject = goco.NSObject{Ptr: goco.InvokeAndReturn("NSMenuItem", "alloc", goco.NSClassFromString("NSMenuItem"))}
	return item
}

func (item NSMenuItem) InitWithTitleActionKeyEquivalent(title foundation.NSString, sel string, keyEq foundation.NSString) NSMenuItem {
	selVar := goco.NSSelectorFromString(sel)
	goco.Invoke("NSMenuItem", "initWithTitle:action:keyEquivalent:", item.Ptr, unsafe.Pointer(&title.Ptr), unsafe.Pointer(&selVar), unsafe.Pointer(&keyEq.Ptr))
	return item
}

func (item NSMenuItem) SetSubMenu(menu NSMenu) {
	goco.Invoke("NSMenuItem", "setSubmenu:", item.Ptr, unsafe.Pointer(&menu.Ptr))
}
