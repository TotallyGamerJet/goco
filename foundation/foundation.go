// +build darwin

package foundation

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation

#import <Foundation/Foundation.h>

void* NSObject_instanceMethodSignatureForSelector(void* obj, _GoString_ sel) {
	return [[((NSObject*)obj) class] instanceMethodSignatureForSelector:NSSelectorFromString([[NSString alloc] initWithBytes:_GoStringPtr(sel) length:_GoStringLen(sel) encoding:NSUTF8StringEncoding])];
}

void* NSObject_methodSignatureForSelector(void* obj, _GoString_ sel) {
	return [((NSObject*)obj) methodSignatureForSelector:NSSelectorFromString([[NSString alloc] initWithBytes:_GoStringPtr(sel) length:_GoStringLen(sel) encoding:NSUTF8StringEncoding])];
}

void* NSObject_performSelectorWithObjectWithObject(void* obj, _GoString_ sel, void* with1, void* with2) {
	return [(NSObject*)obj performSelector:NSSelectorFromString([[NSString alloc] initWithBytes:_GoStringPtr(sel) length:_GoStringLen(sel) encoding:NSUTF8StringEncoding])
			withObject: (NSObject*)with1 withObject: (NSObject*)with2];
}

void* NSObject_performSelectorWithObjectWithObjectClass(_GoString_ cls, _GoString_ sel, void* with1, void* with2) {
	//https://stackoverflow.com/questions/20400366/dynamically-call-static-method-on-class-from-string
	return [NSClassFromString([[NSString alloc] initWithBytes:_GoStringPtr(cls) length:_GoStringLen(cls) encoding:NSUTF8StringEncoding])
				performSelector:NSSelectorFromString([[NSString alloc] initWithBytes:_GoStringPtr(sel) length:_GoStringLen(sel) encoding:NSUTF8StringEncoding])
				withObject: (NSObject*)with1 withObject: (NSObject*)with2];
}

void* NSString_new(_GoString_ str) {
	return [[NSString alloc] initWithBytes:_GoStringPtr(str) length:_GoStringLen(str) encoding:NSUTF8StringEncoding];
}

void* runtime_NSSelectorFromString(void* str) {
	SEL sel = NSSelectorFromString((NSString*)str);
	return sel;
}

*/
import "C"
import (
	"unsafe"
)

type NSObject struct {
	Ptr unsafe.Pointer
}

func (obj NSObject) InstanceMethodSignatureForSelector(selector string) NSMethodSignature {
	return NSMethodSignature{NSObject{C.NSObject_instanceMethodSignatureForSelector(obj.Ptr, selector)}}
}

func (obj NSObject) MethodSignatureForSelector(selector string) NSMethodSignature {
	return NSMethodSignature{NSObject{C.NSObject_methodSignatureForSelector(obj.Ptr, selector)}}
}

func (obj NSObject) PerformSelector(sel string) NSObject {
	return NSObject{C.NSObject_performSelectorWithObjectWithObject(obj.Ptr, sel, nil, nil)}
}

func (obj NSObject) PerformSelectorWithObject(sel string, withObj NSObject) NSObject {
	return NSObject{C.NSObject_performSelectorWithObjectWithObject(obj.Ptr, sel, withObj.Ptr, nil)}
}

func (obj NSObject) PerformSelectorWithObjectWithObject(sel string, withObj, withObj2 NSObject) NSObject {
	return NSObject{C.NSObject_performSelectorWithObjectWithObject(obj.Ptr, sel, withObj.Ptr, withObj2.Ptr)}
}

//Should only be used in dire needs
func PerformSelectorClass(class, selector string) NSObject {
	return NSObject{Ptr: C.NSObject_performSelectorWithObjectWithObjectClass(class, selector, nil, nil)}
}

//Should only be used in dire needs
func PerformSelectorWithObjectClass(class, selector string, obj NSObject) NSObject {
	return NSObject{Ptr: C.NSObject_performSelectorWithObjectWithObjectClass(class, selector, obj.Ptr, nil)}
}

//Should only be used in dire needs
func PerformSelectorWithObjectWithObjectClass(class, selector string, obj, obj2 NSObject) NSObject {
	return NSObject{Ptr: C.NSObject_performSelectorWithObjectWithObjectClass(class, selector, obj.Ptr, obj2.Ptr)}
}

type NSMethodSignature struct {
	NSObject
}

type NSString struct {
	NSObject
}

func (str NSString) Alloc() NSString {
	str.NSObject = PerformSelectorClass("NSString", "alloc")
	return str
}

func (str NSString) Init() NSString {
	str.PerformSelector("init")
	return str
}

func (str NSString) InitWithBytesLengthEncoding(bytes []byte, len int, encoding NSStringEncoding) NSString {
	//- initWithBytes:length:encoding: //TODO: implement
	selector := "initWithBytes:length:encoding:"
	sig := str.MethodSignatureForSelector(selector)
	inv := InvocationWithMethodSignature(sig)
	inv.SetSelector(selector)
	inv.SetArgumentAtIndex(unsafe.Pointer(&bytes[0]), 2)
	inv.SetArgumentAtIndex(unsafe.Pointer(&len), 3)
	inv.SetArgumentAtIndex(unsafe.Pointer(&encoding), 4)
	inv.InvokeWithTarget(str.NSObject)
	return str
}

//Go takes a Go string and converts it to a NSString
func (str NSString) Go(goString string) NSString {
	return str.InitWithBytesLengthEncoding([]byte(goString), len(goString), NSUTF8StringEncoding)
}

func (str NSString) InitWithString(aString NSString) NSString {
	str.PerformSelectorWithObject("init", aString.NSObject)
	return str
}

func NewNSString(str string) NSString {
	return NSString{NSObject{C.NSString_new(str)}}
}

func NSSelectorFromString(str NSString) unsafe.Pointer {
	return C.runtime_NSSelectorFromString(str.Ptr) //C.NSSelectorFromString((*C.void)(str.Ptr))
}
