package appkit

//NSApplicationActivationPolicy Activation policies (used by activationPolicy) that control whether and how an app may be activated.
//
//Reference: https://developer.apple.com/documentation/appkit/nsapplicationactivationpolicy?language=objc
type NSApplicationActivationPolicy int

const (
	//The application is an ordinary app that appears in the Dock and may have a user interface.
	NSApplicationActivationPolicyRegular NSApplicationActivationPolicy = iota
	//The application doesn’t appear in the Dock and doesn’t have a menu bar, but it may be activated programmatically or by clicking on one of its windows.
	NSApplicationActivationPolicyAccessory
	//The application doesn’t appear in the Dock and may not create windows or be activated.
	NSApplicationActivationPolicyProhibited
)

// NSWindowStyleMask These constants specify the style of a window, and can be combined using the C bitwise OR operator.
//
// Reference: https://developer.apple.com/documentation/appkit/nswindowstylemask?language=objc
type NSWindowStyleMask int

const (
	//The window displays none of the usual peripheral elements. Useful only for display or caching purposes. A window that uses NSWindowStyleMaskBorderless can’t become key or main, unless the value of canBecomeKeyWindow or canBecomeMainWindow is YES.
	NSWindowStyleMaskBorderless NSWindowStyleMask = 0
	//The window displays a title bar.
	NSWindowStyleMaskTitled NSWindowStyleMask = 1 << 0
	//The window displays a close button.
	NSWindowStyleMaskClosable NSWindowStyleMask = 1 << 1
	//The window displays a minimize button.
	NSWindowStyleMaskMiniaturizable NSWindowStyleMask = 1 << 2
	//The window can be resized by the user.
	NSWindowStyleMaskResizable NSWindowStyleMask = 1 << 3
	//This constant has no effect, because all windows that include a toolbar use the unified style.
	NSWindowStyleMaskUnifiedTitleAndToolbar NSWindowStyleMask = 1 << 12
	//The window can appear full screen. A fullscreen window does not draw its title bar, and may have special handling for its toolbar. (This mask is automatically toggled when toggleFullScreen: is called.)
	NSWindowStyleMaskFullScreen NSWindowStyleMask = 1 << 14
	//When set, the window’s contentView consumes the full size of the window. Although you can combine this constant with other window style masks, it is respected only for windows with a title bar.
	//Note that using this mask opts in to layer-backing. Use the contentLayoutRect or the contentLayoutGuide to lay out views underneath the title bar–toolbar area.
	NSWindowStyleMaskFullSizeContentView NSWindowStyleMask = 1 << 15
	//The window is a panel or a subclass of NSPanel.
	NSWindowStyleMaskUtilityWindow NSWindowStyleMask = 1 << 4
	//The window is a document-modal panel (or a subclass of NSPanel).
	NSWindowStyleMaskDocModalWindow NSWindowStyleMask = 1 << 6
	//The window is a HUD panel.
	NSWindowStyleMaskHUDWindow NSWindowStyleMask = 1 << 13
)

//NSBackingStoreType These constants specify how the drawing done in a window is buffered by the window device.
//
// Reference: https://developer.apple.com/documentation/appkit/nsbackingstoretype?language=objc
type NSBackingStoreType int

const (
	//The window renders all drawing into a display buffer and then flushes it to the screen.
	NSBackingStoreBuffered NSBackingStoreType = 2
)
