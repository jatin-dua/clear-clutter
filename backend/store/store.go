package store

import (
	"log"

	adb "github.com/zach-klippenstein/goadb"
)

type Data struct {
	Applications *[]string
	Device       *adb.Device
	DeviceInfo   *adb.DeviceInfo
}

var data *Data

func SetTemplateData(apps *[]string, device *adb.Device) {
	deviceInfo, err := device.DeviceInfo()
	if err != nil {
		log.Fatal(err)
	}

	data = &Data{
		Applications: apps,
		Device:       device,
		DeviceInfo:   deviceInfo,
	}
}

func TemplateData() *Data {
	return data
}
