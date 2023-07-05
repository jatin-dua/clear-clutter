package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jatin-dua/clear-clutter/backend/router"
	adb "github.com/zach-klippenstein/goadb"
)

func main() {
	port := flag.Int("p", adb.AdbPort, "")
	flag.Parse()

	client, err := adb.NewWithConfig(adb.ServerConfig{
		Port: *port,
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

	router.Setup()
}
