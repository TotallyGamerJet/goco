package main

import (
	. "goco"
)

//Inspired by this: https://gist.github.com/cstrahan/8cf0002a8b6474f3f5f58e400eec16cf
func main() {
	app := NSApplication{}.SharedApplication()
	app.SetActivationPolicy(NSApplicationActivationPolicyRegular)

	menubar := NSMenu{}.Alloc()
	appMenuItem := NSMenuItem{}.Alloc()
	menubar.AddItem(appMenuItem)
	app.SetMainMenu(menubar)

	appMenu := NSMenu{}.Alloc()
	quitMenuItem := NSMenuItem{}.Alloc().InitWithTitleActionKeyEquivalent(NSString{}.Go("Quit"), "terminate:", NSString{}.Go("q"))
	appMenu.AddItem(quitMenuItem)
	appMenuItem.SetSubMenu(appMenu)

	win := NSWindow{}.Alloc().InitWithContentRectStyleMaskBackingDefer(NSMakeRect(0, 0, 200, 200), NSWindowStyleMaskTitled, NSBackingStoreBuffered, true)
	win.CascadeTopLeftFromPoint(NSMakePoint(20, 20))
	win.SetTitle(NSString{}.Go("Here I am"))
	win.MakeKeyAndOrderFront()

	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
