package goco

import (
	"unsafe"
)

type NSObject struct {
	Ptr unsafe.Pointer
}

func (obj NSObject) Release() {
	Invoke("NSObject", "release", obj.Ptr)
}

func (obj NSObject) Retain() {
	Invoke("NSObject", "retain", obj.Ptr)
}
