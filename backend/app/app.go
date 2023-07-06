package app

import (
	"fmt"
	"log"
	"strings"

	adb "github.com/zach-klippenstein/goadb"
)

func InstalledApps(device adb.Device) []string {
	var installedApps []string
	apps, err := device.RunCommand("pm list packages")
	if err != nil {
		log.Fatal(err)
	}

	appList := strings.Split(apps, "\n")

	for _, app := range appList {
		trimmed := strings.TrimPrefix(app, "package:")
		trimmed = strings.TrimSpace(trimmed)
		if trimmed == "" {
			continue
		}
		installedApps = append(installedApps, trimmed)
	}

	return installedApps
}

func DisplayAllApps(appList []string) {
	for _, name := range appList {
		fmt.Println(name)
	}
}
