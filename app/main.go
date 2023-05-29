package app

import "github.com/robin-samuel/androidgo"

type Application struct {
	packageName string
	device      androidgo.Device
}

func NewApplication(packageName string, device androidgo.Device) Application {
	return Application{packageName: packageName, device: device}
}

func (a *Application) GetPackageName() string {
	return a.packageName
}

func (a *Application) GetDevice() androidgo.Device {
	return a.device
}
