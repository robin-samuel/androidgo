package app

import "fmt"

// ActivityLifecycleCallbacks represents the callbacks for activity lifecycle events.
type ActivityLifecycleCallbacks interface {
	OnActivityCreated(activity Activity)
	OnActivityStarted(activity Activity)
	OnActivityResumed(activity Activity)
	OnActivityPaused(activity Activity)
	OnActivityStopped(activity Activity)
	OnActivityDestroyed(activity Activity)
}

// RegisterActivityLifecycleCallbacks registers an activity lifecycle callback.
func (app *Application) RegisterActivityLifecycleCallbacks(callback ActivityLifecycleCallbacks) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.activityLifecycleCbs = append(app.activityLifecycleCbs, callback)
}

// AddActivity adds an activity to the application.
func (app *Application) AddActivity(activity *Activity) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.Activities = append(app.Activities, activity)
}

// PerformActivityLifecycleCallbacks invokes the registered activity lifecycle callbacks for a given activity and event.
func (app *Application) PerformActivityLifecycleCallbacks(activity Activity, event string) {
	app.mu.Lock()
	defer app.mu.Unlock()

	for _, callback := range app.activityLifecycleCbs {
		switch event {
		case "Created":
			callback.OnActivityCreated(activity)
		case "Started":
			callback.OnActivityStarted(activity)
		case "Resumed":
			callback.OnActivityResumed(activity)
		case "Paused":
			callback.OnActivityPaused(activity)
		case "Stopped":
			callback.OnActivityStopped(activity)
		case "Destroyed":
			callback.OnActivityDestroyed(activity)
		default:
			fmt.Println("Unknown event:", event)
		}
	}
}
