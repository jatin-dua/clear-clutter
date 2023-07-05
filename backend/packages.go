package main

import (
	"fmt"
	"log"
	"strings"

	adb "github.com/zach-klippenstein/goadb"
)

func getInstalledPackages(device adb.Device) []string {
	packages, err := device.RunCommand("pm list packages")
	if err != nil {
		log.Fatal(err)
	}

	packageList := strings.Split(packages, "\n")
	return packageList
}

func displayAllPackages(packageList []string) {
	for _, name := range packageList {
		formatted_name := strings.Split(name, ":")
		if len(formatted_name) >= 2 {
			fmt.Println(formatted_name[1])
		}
	}
}
