// +build darwin

package appkit

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework AppKit

#import <AppKit/AppKit.h>

void* NSMenuItem_initWithTitleactionkeyEquivalent(void* item, void* title, void* action, void* keyEq) {
	return [(NSMenuItem*)item initWithTitle:(NSString*)title action:action keyEquivalent:(NSString*)keyEq];
}


*/
import "C"
import (
	"goco/foundation"
)

type NSMenu struct {
	foundation.NSObject
}

func (m NSMenu) Alloc() NSMenu {
	m.NSObject = foundation.PerformSelectorClass("NSMenu", "alloc")
	return m
}

func (m NSMenu) InitWithTitle(title foundation.NSString) NSMenu {
	m.PerformSelectorWithObject("initWithTitle:", title.NSObject) //TODO: test to make sure this works
	return m
}

func (m NSMenu) AddItem(item NSMenuItem) {
	m.PerformSelectorWithObject("addItem:", item.NSObject)
}

type NSMenuItem struct {
	foundation.NSObject
}

func (item NSMenuItem) Alloc() NSMenuItem {
	item.NSObject = foundation.PerformSelectorClass("NSMenuItem", "alloc")
	return item
}

func (item NSMenuItem) InitWithTitleActionKeyEquivalent(title, action, keyEq foundation.NSString) NSMenuItem {
	/*selector := "initWithTitle:action:keyEquivalent:"
	sig := item.MethodSignatureForSelector(selector)
	inv := foundation.InvocationWithMethodSignature(sig)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&title.Ptr), 2)
	inv.SetArgumentAtIndex(foundation.NSSelectorFromString(action), 3)
	inv.SetArgumentAtIndex(unsafe.Pointer(&keyEq.Ptr), 4)
	inv.InvokeWithTarget(item.NSObject)
	//item.NSObject = foundation.NSObject{Ptr: inv.GetReturnValue()}
	return item*/
	return NSMenuItem{foundation.NSObject{Ptr: C.NSMenuItem_initWithTitleactionkeyEquivalent(item.Ptr, title.Ptr, foundation.NSSelectorFromString(action), keyEq.Ptr)}}
}

func (item NSMenuItem) SetSubMenu(menu NSMenu) {
	item.PerformSelectorWithObject("setSubmenu:", menu.NSObject)
}

type NSApplicationDelegate interface {
	//TODO:
}
