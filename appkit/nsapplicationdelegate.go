package appkit

import "goco/foundation"

// NSApplicationDelegate contains all the methods that can be implemented
type NSApplicationDelegate struct {
	ApplicationDidFinishLaunching func(notification foundation.NSNotification)
}
