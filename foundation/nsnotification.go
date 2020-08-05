package foundation

import "goco"

type NSNotificationName NSString

type NSNotification struct {
	goco.NSObject
}

func (n NSNotification) Name() NSNotificationName {
	return NSNotificationName{NSObject: goco.NSObject{Ptr: goco.InvokeAndReturn("NSNotification", "name", n.Ptr)}}
}

func (n NSNotification) Object() goco.NSObject {
	return goco.NSObject{Ptr: goco.InvokeAndReturn("NSNotification", "object", n.Ptr)}
}
