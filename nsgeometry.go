package goco

import "fmt"

/*
#import <Foundation/NSGeometry.h>
*/

var (
	//An NSPoint structure with both x and y coordinates set to 0.
	NSZeroPoint = NSPoint{}
)

//TODO: NSPoint methods
type NSPoint struct {
	x, y float64
}

func NSMakePoint(x, y float64) NSPoint {
	return NSPoint{x, y}
}

//Returns a Boolean value that indicates whether two points are equal.
func NSEqualPoint(aPoint, bPoint NSPoint) bool {
	return aPoint.x == bPoint.x && aPoint.y == bPoint.y
}

func (point NSPoint) String() string {
	return fmt.Sprintf("{%f,%f}", point.x, point.y)
}

//TODO: NSRect methods
type NSRect struct {
	x, y, width, height float64
}

func NSMakeRect(x, y, width, height float64) NSRect {
	return NSRect{x, y, width, height}
}
