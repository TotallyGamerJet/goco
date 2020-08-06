package foundation

import "goco"

type NSString struct {
	goco.NSObject
}

func (str NSString) Go(goStr string) NSString { //TODO: figure out an actual implementation
	/*bytes := []byte(goStr)
	data := NSData{}.Alloc().DataWithBytesLength(bytes, len(bytes))
	return str.Alloc().InitWithDataEncoding(data, NSUTF8StringEncoding)*/
	return NSString{goco.NSObject{Ptr: goco.NSStringNew(goStr)}}
}

//Impossible because: https://stackoverflow.com/questions/27630390/calls-to-instancemethodsignatureforselector-return-nil-why
/*func (str NSString) InitWithBytesLengthEncoding(b []byte, length int, encoding NSStringEncoding) NSString {
	if length > len(b) {
		panic("not enough bytes")
	}
	if length < 0 {
		panic("unexpected negative length")
	}
	selector := "initWithBytes:length:encoding:"
	inv := InvocationWithMethodSignature("NSPlaceholderString", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&b[0]), 2)
	inv.SetArgumentAtIndex(unsafe.Pointer(&length), 3)
	inv.SetArgumentAtIndex(unsafe.Pointer(&encoding), 4)
	inv.InvokeWithTarget(str.ptr)
	str.NSObject = NSObject{inv.GetReturnValue()}
	return str
}*/
