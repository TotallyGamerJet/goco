package foundation

/*
#import <Foundation/NSGeometry.h>
*/

//TODO: NSPoint methods
type NSPoint struct {
	x, y float64
}

func NSMakePoint(x, y float64) NSPoint {
	return NSPoint{x, y} //Point: C.NSMakePoint(C.double(x), C.double(y))}
}

//TODO: NSRect methods
type NSRect struct {
	x, y, width, height float64
}

func NSMakeRect(x, y, width, height float64) NSRect {
	return NSRect{x, y, width, height} //Rect: C.NSMakeRect(C.double(x), C.double(y), C.double(width), C.double(height))}
}
