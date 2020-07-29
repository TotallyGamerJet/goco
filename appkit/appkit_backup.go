// +build ignore

package appkit

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework AppKit

#import <AppKit/AppKit.h>

void* NSApplication_sharedApplication() {
	return [NSApplication sharedApplication];
}

int NSApplication_setActivationPolicy(int policy) {
	BOOL ret = [NSApp setActivationPolicy:policy];
	return ret ? 1 : 0;
}

void NSApplication_activateIgnoringOtherApps(unsigned char flag) {
	[NSApp activateIgnoringOtherApps:flag];
}

void NSApplication_setMainMenu(void* menu) {
	[NSApp setMainMenu:menu];
}

void NSApplication_run() {
	[NSApp run];
}

void* NSWindow_initWithContentRectstyleMaskbackingdefer(int x, int y, int width, int height, int mask, int storetype, unsigned char defer) {
	return [[NSWindow alloc] initWithContentRect:NSMakeRect(x, y, width, height)
        styleMask:mask backing:storetype defer:defer];
}

void NSWindow_cascadeTopLeftFromPoint(void* window, int x, int y) {
	[((NSWindow*)window) cascadeTopLeftFromPoint:NSMakePoint(x, y)];
}

void NSWindow_setTitle(void* window, _GoString_ source) {
	[((NSWindow*)window) setTitle:[[NSString alloc] initWithBytes:_GoStringPtr(source) length:_GoStringLen(source) encoding:NSUTF8StringEncoding]];
}

void NSWindow_makeKeyAndOrderFront(void* window) {
	[((NSWindow*)window) makeKeyAndOrderFront:nil];
}

void* NSMenu_new() {
	return [NSMenu alloc];
}

void NSMenu_addItem(void* menu, void* item) {
	[((NSMenu*)menu) addItem:((NSMenuItem*)item)];
}

void* NSMenuItem_new() {
	return [NSMenuItem alloc];
}

void* NSMenuItem_initWithTitleactionkeyEquivalent(_GoString_ title, _GoString_ action, _GoString_ keyEq) {
	return [[NSMenuItem alloc] initWithTitle:[[NSString alloc] initWithBytes:_GoStringPtr(title) length:_GoStringLen(title) encoding:NSUTF8StringEncoding]
        action:NSSelectorFromString([[NSString alloc] initWithBytes:_GoStringPtr(action) length:_GoStringLen(action) encoding:NSUTF8StringEncoding])
		keyEquivalent:[[NSString alloc] initWithBytes:_GoStringPtr(keyEq) length:_GoStringLen(keyEq) encoding:NSUTF8StringEncoding]];
}

void NSMenuItem_setSubMenu(void* item, void* menu) {
	[(NSMenuItem*)item setSubmenu:(NSMenu*)menu];
}

*/
import "C"
import (
	"goco"
	"unsafe"
)

type NSApplication struct {
	ptr unsafe.Pointer
}

func SharedApplication() NSApplication {
	return NSApplication{ptr: C.NSApplication_sharedApplication()}
}

func (app NSApplication) SetActivationPolicy(policy NSApplicationActivationPolicy) bool {
	var success = C.NSApplication_setActivationPolicy(C.int(policy))
	return goco.IntToBool(int(success))
}

func (app NSApplication) ActivateIgnoringOtherApps(flag bool) {
	C.NSApplication_activateIgnoringOtherApps(C.uchar(goco.BoolToBOOL(flag)))
}

func (app NSApplication) SetMainMenu(menu NSMenu) {
	C.NSApplication_setMainMenu(menu.ptr)
}

func (app NSApplication) Run() {
	//C.NSApplication_run()
	C.NSObject_performSelector(app.ptr, "run")
}

type NSWindow struct {
	ptr unsafe.Pointer
}

func InitWindow(x, y, width, height int, styleMask NSWindowStyleMask, storeType NSBackingStoreType, doDefer bool) NSWindow { //todo: utilize NSRect in Foundation
	return NSWindow{C.NSWindow_initWithContentRectstyleMaskbackingdefer(C.int(x), C.int(y), C.int(width), C.int(height), C.int(styleMask), C.int(storeType), C.uchar(goco.BoolToBOOL(doDefer)))}
}

func (win NSWindow) CascadeTopLeftFromPoint(x, y int) { //TODO: utilize NSPoint in Foundation
	//TODO: get return NSPoint
	C.NSWindow_cascadeTopLeftFromPoint(win.ptr, C.int(x), C.int(y))
}

func (win NSWindow) SetTitle(title string) {
	C.NSWindow_setTitle(win.ptr, title)
}

func (win NSWindow) MakeKeyAndOrderFront() { //TODO: taken in sender id
	C.NSWindow_makeKeyAndOrderFront(win.ptr)
}

type NSMenu struct {
	ptr unsafe.Pointer
}

func NewMenu() NSMenu {
	return NSMenu{ptr: C.NSMenu_new()}
}

func (m NSMenu) AddItem(item NSMenuItem) {
	C.NSMenu_addItem(m.ptr, item.ptr)
}

type NSMenuItem struct {
	ptr unsafe.Pointer
}

func NewMenuItem() NSMenuItem {
	return NSMenuItem{ptr: C.NSMenuItem_new()}
}

//Recall that a colon (“:”) is part of a method name; setHeight is not the same as setHeight:.
func InitMenuItem(title, action, keyEquivalent string) NSMenuItem {
	return NSMenuItem{ptr: C.NSMenuItem_initWithTitleactionkeyEquivalent(title, action, keyEquivalent)}
}

func (item NSMenuItem) SetSubMenu(menu NSMenu) {
	C.NSMenuItem_setSubMenu(item.ptr, menu.ptr)
}

type NSApplicationDelegate interface {
	//TODO:
}
