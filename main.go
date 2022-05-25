package main

import (
	"fmt"
	"go-nm/networkmanager"
)

func main() {

	savedconnections, err := networkmanager.ListSavedConnections()
	if err != nil {
		panic(err)
	}
	fmt.Println("SETTINGS", savedconnections)

	add, err := networkmanager.AddNetwork("test", "PASSWORD")
	if err != nil {
		panic(err)
	}
	fmt.Println("ADDED NETWORK", add)

	connectivity, err := networkmanager.CheckConnectivity()
	if err != nil {
		panic(err)
	}
	fmt.Println("CONNECTIVITY", connectivity)

	settings, err := networkmanager.GetConnectionSettings("/org/freedesktop/NetworkManager/Settings/71")

	if err != nil {
		panic(err)
	}
	fmt.Println("SETTINGS", settings)

	aps, err := networkmanager.GetAccessPoints("/org/freedesktop/NetworkManager/Devices/2")

	if err != nil {
		panic(err)
	}

	fmt.Println("ACCESS POINTS", aps)

	devices, err := networkmanager.ListDevices()

	if err != nil {
		panic(err)
	}

	fmt.Println("DEVICES", devices)

	wireless, err := networkmanager.GetDeviceByType(2)
	if err != nil {
		fmt.Println(err)
	}
	// Get path of wireless

	fmt.Println("WIRELESS", wireless)

	for _, attr := range wireless {
		println(attr.(map[string]interface{})["State"].(uint32))
	}

	wired, err := networkmanager.GetDeviceByType(1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("WIRED", wired)

}
