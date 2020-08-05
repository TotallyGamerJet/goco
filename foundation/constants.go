package foundation

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
