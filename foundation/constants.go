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
	//TODO: the rest
)
