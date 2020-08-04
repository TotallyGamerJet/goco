// +build ignore

package goco

import "unsafe"

type NSData struct {
	NSObject
}

func (data NSData) Alloc() NSData {
	selector := "alloc"
	inv := InvocationWithMethodSignature("NSData", selector)
	inv.SetSelector(selector)
	inv.InvokeWithTarget(NSClassFromString("NSData"))
	data.NSObject = NSObject{inv.GetReturnValue()}
	return data
}

func (data NSData) DataWithBytesLength(bytes []byte, length int) NSData {
	if length > len(bytes) {
		panic("not enough bytes")
	}
	selector := "dataWithBytes:length:"
	inv := InvocationWithMethodSignature("NSData", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&bytes[0]), 2)
	inv.SetArgumentAtIndex(unsafe.Pointer(&length), 3)
	inv.InvokeWithTarget(data.ptr)
	return data
}
