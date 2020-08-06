package plugin

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#import <Foundation/NSInvocation.h>
#import <Foundation/NSString.h>

void* NSInvocation_invocationWithMethodSignature(_GoString_ obj, _GoString_ selector) {
	NSString* clsStr = [[NSString alloc] initWithBytes:_GoStringPtr(obj) length:_GoStringLen(obj) encoding:NSUTF8StringEncoding];
	Class cls = NSClassFromString(clsStr);
	if (cls == nil) {
		NSLog(@"class is nil %@", [[NSString alloc] initWithBytes:_GoStringPtr(obj) length:_GoStringLen(obj) encoding:NSUTF8StringEncoding]);
	}
	NSString* selStr = [[NSString alloc] initWithBytes:_GoStringPtr(selector) length:_GoStringLen(selector) encoding:NSUTF8StringEncoding];
	SEL sel = NSSelectorFromString(selStr);
	if (sel == nil) {
		NSLog(@"sel is nil %@", [[NSString alloc] initWithBytes:_GoStringPtr(selector) length:_GoStringLen(selector) encoding:NSUTF8StringEncoding]);
	}
	NSMethodSignature* sig = [cls methodSignatureForSelector:sel];
	if (sig == nil) {
		sig = [cls instanceMethodSignatureForSelector:sel];
		if (sig == nil) {
			NSLog(@"could not find method signature for %@ %@", clsStr, selStr);
		}
	}
	return [NSInvocation invocationWithMethodSignature:sig];
}

void NSInvocation_setTarget(void* inv, void* target) {
	[(NSInvocation*)inv setTarget:target];
}

void NSInvocation_setSelector(void* inv, _GoString_ sel) {
	[(NSInvocation*)inv setSelector:NSSelectorFromString([[NSString alloc] initWithBytes:_GoStringPtr(sel) length:_GoStringLen(sel) encoding:NSUTF8StringEncoding])];
}

void NSInvocation_setArgumentatIndex(void* inv, void* arg, int index) {
	[(NSInvocation*)inv setArgument:arg atIndex:index];
}

void NSInvocation_invoke(void* inv) {
	[(NSInvocation*)inv invoke];
}

void NSInvocation_invokeWithTarget(void* inv, void* target) {
	[(NSInvocation*)inv invokeWithTarget: target];
}

void* NSInvocation_getReturnValue(void* inv) {
	void* tempResult;//https://stackoverflow.com/questions/22018272/nsinvocation-returns-value-but-makes-app-crash-with-exc-bad-access
	[(NSInvocation*)inv getReturnValue:&tempResult];
	void* out = (__bridge void*)tempResult;
	return out;
}

void* runtime_NSSelectorFromString(_GoString_ str) {
	return NSSelectorFromString([[[NSString alloc] initWithBytes:_GoStringPtr(str) length:_GoStringLen(str) encoding:NSUTF8StringEncoding] autorelease]);
}

void* runtime_NSClassFromString(_GoString_ str) {
	return NSClassFromString([[[NSString alloc] initWithBytes:_GoStringPtr(str) length:_GoStringLen(str) encoding:NSUTF8StringEncoding] autorelease]);
}

void* NSString_goString(_GoString_ str) {
	return [[NSString alloc] initWithBytes:_GoStringPtr(str) length:_GoStringLen(str) encoding:NSUTF8StringEncoding];
}

void* NSInvocation_invokeAndReturn(_GoString_ className, _GoString_ selector, void* target, void* a, void* b, void* c) {
	NSString* clsStr = [[NSString alloc] initWithBytes:_GoStringPtr(className) length:_GoStringLen(className) encoding:NSUTF8StringEncoding];
	Class cls = NSClassFromString(clsStr);
	if (cls == nil) {
		NSLog(@"class is nil %@", clsStr);
	}
	NSString* selStr = [[NSString alloc] initWithBytes:_GoStringPtr(selector) length:_GoStringLen(selector) encoding:NSUTF8StringEncoding];
	SEL sel = NSSelectorFromString(selStr);
	if (sel == nil) {
		NSLog(@"sel is nil %@", selStr);
	}
	NSMethodSignature* sig = [cls methodSignatureForSelector:sel];
	if (sig == nil) {
		sig = [cls instanceMethodSignatureForSelector:sel];
		if (sig == nil) {
			NSLog(@"could not find method signature for %@ %@", clsStr, selStr);
		}
	}
	NSInvocation* inv =  [NSInvocation invocationWithMethodSignature:sig];
	[inv setSelector:sel];
	if (a != nil) {
		[inv setArgument:a atIndex:2];
		if(b != nil) {
			[inv setArgument:b atIndex:3];
			if (c != nil) {
				[inv setArgument:c atIndex:4];
			}
		}
	}
	[inv invokeWithTarget:target];
	void* tempResult;
	[inv getReturnValue:&tempResult];
	void* out = (__bridge void*)tempResult;
	return out;
}

void NSInvocation_justInvoke(_GoString_ className, _GoString_ selector,void* target, void* a, void* b, void* c) {
	NSString* clsStr = [[NSString alloc] initWithBytes:_GoStringPtr(className) length:_GoStringLen(className) encoding:NSUTF8StringEncoding];
	Class cls = NSClassFromString(clsStr);
	if (cls == nil) {
		NSLog(@"class is nil %@", clsStr);
	}
	NSString* selStr = [[NSString alloc] initWithBytes:_GoStringPtr(selector) length:_GoStringLen(selector) encoding:NSUTF8StringEncoding];
	SEL sel = NSSelectorFromString(selStr);
	if (sel == nil) {
		NSLog(@"sel is nil %@", selStr);
	}
	NSMethodSignature* sig = [cls methodSignatureForSelector:sel];
	if (sig == nil) {
		sig = [cls instanceMethodSignatureForSelector:sel];
		if (sig == nil) {
			NSLog(@"could not find method signature for %@ %@", clsStr, selStr);
		}
	}
	NSInvocation* inv =  [NSInvocation invocationWithMethodSignature:sig];
	[inv setSelector:sel];
	if (a != nil) {
		[inv setArgument:a atIndex:2];
		if(b != nil) {
			[inv setArgument:b atIndex:3];
			if (c != nil) {
				[inv setArgument:c atIndex:4];
			}
		}
	}
	[inv invokeWithTarget:target];
}

*/
import "C"
import (
	"unsafe"
)

//NSSelectorFromString gets a selector for the string. DON'T take the address of any variables
//that you assign to this.
func NSSelectorFromString(str string) unsafe.Pointer {
	return C.runtime_NSSelectorFromString(str)
}

func NSClassFromString(str string) unsafe.Pointer {
	return C.runtime_NSClassFromString(str)
}

type NSInvocation struct {
	Ptr unsafe.Pointer
}

func InvocationWithMethodSignature(objectName, selectorName string) NSInvocation {
	return NSInvocation{C.NSInvocation_invocationWithMethodSignature(objectName, selectorName)}
}

func (inv NSInvocation) SetArgumentAtIndex(arg unsafe.Pointer, index int) {
	C.NSInvocation_setArgumentatIndex(inv.Ptr, arg, C.int(index))
}

func (inv NSInvocation) SetSelector(sel string) {
	C.NSInvocation_setSelector(inv.Ptr, sel)
}

func (inv NSInvocation) SetTarget(target unsafe.Pointer) {
	if target == nil {
		panic("target is nil")
	}
	C.NSInvocation_setTarget(inv.Ptr, target)
}

func (inv NSInvocation) Invoke() {
	C.NSInvocation_invoke(inv.Ptr)
}

func (inv NSInvocation) InvokeWithTarget(target unsafe.Pointer) {
	if target == nil {
		panic("target is nil")
	}
	C.NSInvocation_invokeWithTarget(inv.Ptr, target)
}

func (inv NSInvocation) GetReturnValue() unsafe.Pointer {
	return C.NSInvocation_getReturnValue(inv.Ptr)
}

func NsstringGostring(str string) unsafe.Pointer {
	return C.NSString_goString(str)
}

func InvokeAndReturn(className, selector string, target unsafe.Pointer, args ...unsafe.Pointer) unsafe.Pointer { //TODO: use args
	if target == nil {
		panic("target is nil")
	}
	if len(args) > 3 {
		panic("more than 3 arguments supplied")
	}
	var a, b, c unsafe.Pointer
	switch len(args) {
	case 3:
		c = args[2]
		fallthrough
	case 2:
		b = args[1]
		fallthrough
	case 1:
		a = args[0]
	}
	return C.NSInvocation_invokeAndReturn(className, selector, target, a, b, c)
}

func Invoke(className, selector string, target unsafe.Pointer, args ...unsafe.Pointer) { //TODO: use args
	if target == nil {
		panic("target is nil")
	}
	if len(args) > 3 {
		panic("more than 3 arguments supplied")
	}
	var a, b, c unsafe.Pointer
	switch len(args) {
	case 3:
		c = args[2]
		fallthrough
	case 2:
		b = args[1]
		fallthrough
	case 1:
		a = args[0]
	}
	C.NSInvocation_justInvoke(className, selector, target, a, b, c)
}

//helpful links
//https://stackoverflow.com/questions/904515/how-to-use-performselectorwithobjectafterdelay-with-primitive-types-in-cocoa
