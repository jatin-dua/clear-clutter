package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jatin-dua/clear-clutter/backend/app"
	"github.com/jatin-dua/clear-clutter/backend/routes"
	"github.com/jatin-dua/clear-clutter/backend/store"
	adb "github.com/zach-klippenstein/goadb"
)

const listenAddr = 8080

func main() {
	adbPort := flag.Int("adb", adb.AdbPort, "Port for ADB server")
	listenAddr := flag.Int("port", listenAddr, "Port for starting Application")
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
	fmt.Printf("ADB v%v\n", serverVersion)

	devices, err := client.ListDevices()
	if err != nil {
		log.Fatal(err)
	}

	if len(devices) == 0 {
		log.Fatal("No device connected.")
	}

	descriptor := adb.AnyDevice()
	device := client.Device(descriptor)

	installedApps := app.InstalledApps(*device)

	store.SetTemplateData(&installedApps, device)
	routes.SetupRoutes(*listenAddr)
}
