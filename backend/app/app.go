package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/jatin-dua/clear-clutter/backend/store"
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

func UninstallApps(appList []string) {
	for _, app := range appList {
		response, err := store.TemplateData().Device.RunCommand(fmt.Sprintf("pm uninstall -k --user 0 %v", app))
		if err != nil {
			log.Println("uninstall error: ", err)
		}

		fmt.Println(app, "Uninstall", response)
	}
}
