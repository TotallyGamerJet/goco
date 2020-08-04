package goco

type NSString struct {
	NSObject
}

/*
func (str NSString) Alloc() NSString {
	selector := "alloc"
	inv := InvocationWithMethodSignature("NSString", selector)
	inv.SetSelector(selector)
	inv.InvokeWithTarget(NSClassFromString("NSString"))
	str.NSObject = NSObject{inv.GetReturnValue()}
	return str
}

func (str NSString) InitWithDataEncoding(data NSData, encoding NSStringEncoding) NSString {
	selector := "initWithData:encoding:"
	inv := InvocationWithMethodSignature("NSString", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(data.ptr, 2)
	inv.SetArgumentAtIndex(unsafe.Pointer(&encoding), 3)
	inv.InvokeWithTarget(str.ptr)
	return str
}*/

func (str NSString) Go(goStr string) NSString { //TODO: figure out an actual implementation
	/*bytes := []byte(goStr)
	data := NSData{}.Alloc().DataWithBytesLength(bytes, len(bytes))
	return str.Alloc().InitWithDataEncoding(data, NSUTF8StringEncoding)*/
	return NSString{NSObject{ptr: NSStringNew(goStr)}}
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
/*
func (str NSString) StringWithUTF8String(goStr string) NSString {
	b := []byte(goStr)
	//c := C.CString(goStr)
	selector := "stringWithUTF8String:"
	inv := InvocationWithMethodSignature("NSString", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&b[0]), 2)
	inv.SetTarget(NSClassFromString("NSString"))
	inv.Invoke()
	str.NSObject = NSObject{inv.GetReturnValue()}
	return str
}*/

/*
func (str NSString) InitWithUTF8String(goStr string) NSString {
	b := append([]byte(goStr), 0)
	selector := "initWithUTF8String:"
	inv := InvocationWithMethodSignature("NSString", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&b[0]), 2)
	if str.ptr == nil {
		panic("nil string")
	}
	inv.InvokeWithTarget(str.ptr)
	return str
}*/
