package goco

import (
	"goco/plugin"
	"unsafe"
)

var NSClassFromString = plugin.NSClassFromString
var NSSelectorFromString = plugin.NSSelectorFromString
var NSStringNew = plugin.NsstringGostring
var InvokeAndReturn = plugin.InvokeAndReturn
var Invoke = plugin.Invoke

type NSInvocation interface {
	SetArgumentAtIndex(arg unsafe.Pointer, index int)
	SetSelector(sel string)
	SetTarget(target unsafe.Pointer)
	Invoke()
	InvokeWithTarget(target unsafe.Pointer)
	GetReturnValue() unsafe.Pointer
}

func InvocationWithMethodSignature(objName, selector string) NSInvocation {
	return plugin.InvocationWithMethodSignature(objName, selector)
}
