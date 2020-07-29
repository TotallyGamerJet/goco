package foundation

/*
#import <Foundation/NSGeometry.h>
*/
import "C"

//TODO: NSPoint methods
type NSPoint struct {
	Point C.NSPoint
}

func NSMakePoint(x, y float64) NSPoint {
	return NSPoint{Point: C.NSMakePoint(C.double(x), C.double(y))}
}

//TODO: NSRect methods
type NSRect struct {
	Rect C.NSRect
}

func NSMakeRect(x, y, width, height float64) NSRect {
	return NSRect{Rect: C.NSMakeRect(C.double(x), C.double(y), C.double(width), C.double(height))}
}
