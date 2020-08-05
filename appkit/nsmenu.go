package appkit

import (
	"goco"
	"unsafe"
)

type NSMenu struct {
	goco.NSObject
}

func (menu NSMenu) Alloc() NSMenu {
	menu.NSObject = goco.NSObject{Ptr: goco.InvokeAndReturn("NSMenu", "alloc", goco.NSClassFromString("NSMenu"))}
	return menu
}

func (menu NSMenu) AddItem(item NSMenuItem) {
	goco.Invoke("NSMenu", "addItem:", menu.Ptr, unsafe.Pointer(&item))
}
