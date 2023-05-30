package app

import (
	"sync"

	"github.com/robin-samuel/androidgo"
)

// Application represents the main application.
type Application struct {
	// Main Properties
	packageName string
	device      androidgo.Device

	// Activity Lifecycle
	Activities           []*Activity
	mu                   sync.Mutex
	activityLifecycleCbs []ActivityLifecycleCallbacks
}

// NewApplication creates a new Application instance.
func NewApplication(packageName string, device androidgo.Device) *Application {
	return &Application{
		packageName: packageName,
		device:      device,

		Activities:           make([]*Activity, 0),
		activityLifecycleCbs: make([]ActivityLifecycleCallbacks, 0),
	}
}

// Return the package name of the application.
func (a *Application) GetPackageName() string {
	return a.packageName
}

// Return the device running the application.
func (a *Application) GetDevice() androidgo.Device {
	return a.device
}
