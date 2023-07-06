package app

import (
	"fmt"
	"log"
	"strings"

	adb "github.com/zach-klippenstein/goadb"
)

func InstalledApps(device adb.Device) []string {
	apps, err := device.RunCommand("pm list packages")
	if err != nil {
		log.Fatal(err)
	}

	appList := strings.Split(apps, "\n")
	return appList
}

func DisplayAllApps(appList []string) {
	for _, name := range appList {
		formatted_name := strings.Split(name, ":")
		if len(formatted_name) >= 2 {
			fmt.Println(formatted_name[1])
		}
	}
}
