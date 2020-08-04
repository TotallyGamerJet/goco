package goco

import (
	"unsafe"
)

type NSObject struct {
	ptr unsafe.Pointer
}

/*
func (obj NSObject) PerformSelector(objName, sel string) NSObject {
	inv := goco.InvocationWithMethodSignature(objName, sel)
	inv.SetSelector(sel)
	inv.InvokeWithTarget(obj.ptr)
	return NSObject{inv.GetReturnValue()}
}*/
