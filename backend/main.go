package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jatin-dua/clear-clutter/backend/app"
	"github.com/jatin-dua/clear-clutter/backend/routes"
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

	appList := app.InstalledApps(*device)
	app.DisplayAllApps(appList)

	routes.SetData(&appList)

	routes.SetupRoutes(PORT)
}
