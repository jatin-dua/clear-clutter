package main

import (
	"flag"
	"fmt"
	"log"

	adb "github.com/zach-klippenstein/goadb"
)

const PORT = 8080

func main() {
	adbPort := flag.Int("p", adb.AdbPort, "")
	flag.Parse()

	client, err := adb.NewWithConfig(adb.ServerConfig{
		Port: *adbPort,
	})
	if err != nil {
		log.Fatal(err)
	}
	client.StartServer()

	serverVersion, err := client.ServerVersion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server version:", serverVersion)

	devices, err := client.ListDevices()
	if err != nil {
		log.Fatal(err)
	}

	if len(devices) == 0 {
		log.Fatal("No device connected.")
	}

	fmt.Println("Devices:")
	for _, device := range devices {
		fmt.Printf("\t%+v\n", *device)
	}

	descriptor := adb.AnyDevice()
	device := client.Device(descriptor)

	packages := getInstalledPackages(*device)
	fmt.Println(packages)

	SetupRoutes(PORT)
}

func getInstalledPackages(device adb.Device) string {
	packages, err := device.RunCommand("pm list packages")
	if err != nil {
		log.Fatal(err)
	}

	return packages
}
