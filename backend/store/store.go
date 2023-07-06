package store

import adb "github.com/zach-klippenstein/goadb"

type Data struct {
	Applications *[]string
	Device       *adb.DeviceInfo
}

var data Data

func SetTemplateData(apps *[]string, device *adb.DeviceInfo) {
	data = Data{
		Applications: apps,
		Device:       device,
	}
}

func TemplateData() *Data {
	return &data
}
