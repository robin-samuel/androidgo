package app

// Activity represents an activity in the application.
type Activity struct {
	LocalClassName string
	SavedInstance  Bundle
}

// GetLocalClassName returns the local class name of the activity.
func (a *Activity) GetLocalClassName() string {
	return a.LocalClassName
}
