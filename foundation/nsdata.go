// +build ignore

package foundation

import (
	"goco"
	"unsafe"
)

type NSData struct {
	goco.NSObject
}

func (data NSData) Alloc() NSData {
	selector := "alloc"
	inv := goco.InvocationWithMethodSignature("NSData", selector)
	inv.SetSelector(selector)
	inv.InvokeWithTarget(goco.NSClassFromString("NSData"))
	data.NSObject = goco.NSObject{inv.GetReturnValue()}
	return data
}

func (data NSData) DataWithBytesLength(bytes []byte, length int) NSData {
	if length > len(bytes) {
		panic("not enough bytes")
	}
	selector := "dataWithBytes:length:"
	inv := goco.InvocationWithMethodSignature("NSData", selector)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&bytes[0]), 2)
	inv.SetArgumentAtIndex(unsafe.Pointer(&length), 3)
	inv.InvokeWithTarget(data.ptr)
	return data
}
