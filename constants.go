package goco

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

// NSStringEncoding The following constants are provided by NSString as possible string encodings.
//
// Reference: https://developer.apple.com/documentation/foundation/nsstringencoding
type NSStringEncoding int

const (
	//Strict 7-bit ASCII encoding within 8-bit chars; ASCII values 0…127 only.
	NSASCIIStringEncoding NSStringEncoding = iota + 1
	//8-bit ASCII encoding with NEXTSTEP extensions.
	NSNEXTSTEPStringEncoding
	//8-bit EUC encoding for Japanese text.
	NSJapaneseEUCStringEncoding
	//An 8-bit representation of Unicode characters, suitable for transmission or storage by ASCII-based systems.
	NSUTF8StringEncoding
	//8-bit ISO Latin 1 encoding.
	NSISOLatin1StringEncoding
	//8-bit Adobe Symbol encoding vector.
	NSSymbolStringEncoding
	//7-bit verbose ASCII to represent all Unicode characters.
	NSNonLossyASCIIStringEncoding
	//8-bit Shift-JIS encoding for Japanese text.
	NSShiftJISStringEncoding
	//8-bit ISO Latin 2 encoding.
	NSISOLatin2StringEncoding
	//The canonical Unicode encoding for string objects.
	NSUnicodeStringEncoding
	//Microsoft Windows codepage 1251, encoding Cyrillic characters; equivalent to AdobeStandardCyrillic font encoding.
	NSWindowsCP1251StringEncoding
	//Microsoft Windows codepage 1252; equivalent to WinLatin1.
	NSWindowsCP1252StringEncoding
	//Microsoft Windows codepage 1253, encoding Greek characters.
	NSWindowsCP1253StringEncoding
	//Microsoft Windows codepage 1254, encoding Turkish characters.
	NSWindowsCP1254StringEncoding
	//Microsoft Windows codepage 1250; equivalent to WinLatin2.
	NSWindowsCP1250StringEncoding
	//ISO 2022 Japanese encoding for email.
	NSISO2022JPStringEncoding NSStringEncoding = 21
	//Classic Macintosh Roman encoding.
	NSMacOSRomanStringEncoding NSStringEncoding = 30
	//An alias for NSUnicodeStringEncoding.
	NSUTF16StringEncoding = NSUnicodeStringEncoding
	//NSUTF16StringEncoding encoding with explicit endianness specified.
	NSUTF16BigEndianStringEncoding NSStringEncoding = 0x90000100
	//NSUTF16StringEncoding encoding with explicit endianness specified.
	NSUTF16LittleEndianStringEncoding NSStringEncoding = 0x94000100
	//32-bit UTF encoding.
	NSUTF32StringEncoding NSStringEncoding = 0x8c000100
	//NSUTF32StringEncoding encoding with explicit endianness specified.
	NSUTF32BigEndianStringEncoding NSStringEncoding = 0x98000100
	//NSUTF32StringEncoding encoding with explicit endianness specified.
	NSUTF32LittleEndianStringEncoding NSStringEncoding = 0x9c000100
)
