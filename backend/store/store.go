package store

import (
	"log"

	"github.com/jatin-dua/clear-clutter/backend/app"
	adb "github.com/zach-klippenstein/goadb"
)

type Data struct {
	Applications *[]string
	Device       *adb.DeviceInfo
}

var data Data
var device *adb.Device

func SetTemplateData(apps *[]string, dev *adb.Device) {
	deviceInfo, err := dev.DeviceInfo()
	if err != nil {
		log.Fatal(err)
	}

	data = Data{
		Applications: apps,
		Device:       deviceInfo,
	}

	device = dev
}

func TemplateData() *Data {
	return &data
}

func RefreshData() {
	installedApps := app.InstalledApps(*device)
	SetTemplateData(&installedApps, device)
}
