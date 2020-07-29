package foundation

/*
#import <Foundation/NSInvocation.h>
#import <Foundation/NSString.h>

void NSInvocation_setTarget(void* inv, void* target) {
	[(NSInvocation*)inv setTarget:target];
}

void NSInvocation_setTargetClass(void* inv, _GoString_ target) {
	[(NSInvocation*)inv setTarget:NSClassFromString([[NSString alloc] initWithBytes:_GoStringPtr(target) length:_GoStringLen(target) encoding:NSUTF8StringEncoding])];
}

void NSInvocation_setSelector(void* inv, _GoString_ sel) {
	[(NSInvocation*)inv setSelector:NSSelectorFromString([[NSString alloc] initWithBytes:_GoStringPtr(sel) length:_GoStringLen(sel) encoding:NSUTF8StringEncoding])];
}

void NSInvocation_setArgumentatIndex(void* inv, void* arg, int index) {
	[(NSInvocation*)inv setArgument:arg atIndex:index];
}

void* NSInvocation_getReturnValue(void* inv) {
	void* tempResult;//https://stackoverflow.com/questions/22018272/nsinvocation-returns-value-but-makes-app-crash-with-exc-bad-access
	[(NSInvocation*)inv getReturnValue:&tempResult];
	void* out = (__bridge void*)tempResult;
	return out;
}

*/
import "C"
import "unsafe"

type NSInvocation struct {
	NSObject
}

func InvocationWithMethodSignature(sig NSMethodSignature) NSInvocation {
	/*void* NSInvocation_invocationWithMethodSignature(void* sig) {
		return [NSInvocation invocationWithMethodSignature:(NSMethodSignature*)sig];
	}
	//NSObject{C.NSInvocation_invocationWithMethodSignature(sig.Ptr)}}
	*/
	return NSInvocation{PerformSelectorWithObjectClass("NSInvocation", "invocationWithMethodSignature:", sig.NSObject)}
}

func (inv NSInvocation) SetArgumentAtIndex(arg unsafe.Pointer, index int) {
	C.NSInvocation_setArgumentatIndex(inv.Ptr, arg, C.int(index))
}

func (inv NSInvocation) SetSelector(sel string) {
	C.NSInvocation_setSelector(inv.Ptr, sel)
}

func (inv NSInvocation) SetTarget(target unsafe.Pointer) {
	//inv.PerformSelectorWithObject("setTarget:", target)
	C.NSInvocation_setTarget(inv.Ptr, target)
}

func (inv NSInvocation) SetTargetClass(target string) {
	C.NSInvocation_setTargetClass(inv.Ptr, target)
}

func (inv NSInvocation) Invoke() {
	inv.PerformSelector("invoke")
}

func (inv NSInvocation) InvokeWithTarget(target NSObject) {
	inv.PerformSelectorWithObject("invokeWithTarget:", target)
}
func (inv NSInvocation) GetReturnValue() unsafe.Pointer {
	return C.NSInvocation_getReturnValue(inv.Ptr)
}
