package main

import (
	. "goco/appkit"
	"goco/foundation"
)

//Example from: https://gist.github.com/cstrahan/8cf0002a8b6474f3f5f58e400eec16cf
func main() {
	app := NSApplication{}.SharedApplication()
	app.SetActivationPolicy(NSApplicationActivationPolicyRegular)

	menubar := NSMenu{}.Alloc()
	appMenuItem := NSMenuItem{}.Alloc()
	menubar.AddItem(appMenuItem)
	app.SetMainMenu(menubar)

	appMenu := NSMenu{}.Alloc()
	quitMenuItem := NSMenuItem{}.Alloc().InitWithTitleActionKeyEquivalent(foundation.NewNSString("Quit"), foundation.NewNSString("terminate:"), foundation.NewNSString("q"))
	appMenu.AddItem(quitMenuItem)
	appMenuItem.SetSubMenu(appMenu)

	win := NSWindow{}.Alloc().InitWithContentRectStyleMaskBackingDefer(foundation.NSMakeRect(0, 0, 200, 200), NSWindowStyleMaskTitled, NSBackingStoreBuffered, true)
	win.CascadeTopLeftFromPoint(foundation.NSMakePoint(20, 20))
	win.SetTitle(foundation.NewNSString("What up dawg!")) //foundation.NSString{}.Alloc().Go("Here we go!")) //
	win.MakeKeyAndOrderFront()

	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
