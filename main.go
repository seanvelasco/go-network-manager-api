package main

import (
	"encoding/json"
	"fmt"
	"go-nm/networkmanager"
	"net/http"

	"github.com/godbus/dbus/v5"
)

var ResponseMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func settings(w http.ResponseWriter, req *http.Request) {
	savedconnections, err := networkmanager.ListSavedConnections()
	if err != nil {
		panic(err)
	}
	jsonString, err := json.Marshal(savedconnections)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", jsonString)

}

func getConnectivity(w http.ResponseWriter, req *http.Request) {
	connectivity, err := networkmanager.CheckConnectivity()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%d", connectivity)
}

func getAllDevices(w http.ResponseWriter, req *http.Request) {
	devices, err := networkmanager.ListDevices()
	if err != nil {
		panic(err)
	}
	jsonString, err := json.Marshal(devices)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", jsonString)
}

func getWiredDevices(w http.ResponseWriter, req *http.Request) {
	devices, err := networkmanager.GetDeviceByType(1)

	if err != nil {
		fmt.Println(err)
	}
	jsonString, err := json.Marshal(devices)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", jsonString)
}

func getWirelessDevices(w http.ResponseWriter, req *http.Request) {
	devices, err := networkmanager.GetDeviceByType(2)

	if err != nil {
		fmt.Println(err)
	}
	jsonString, err := json.Marshal(devices)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", jsonString)
}

func addNetwork(w http.ResponseWriter, req *http.Request) {

	type network struct {
		Ssid     string `json:"ssid"`
		Password string `json:"password"`
	}

	var n network

	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&n)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ResponseMessage.Success = false
		ResponseMessage.Message = "Failed to add" + n.Ssid
		jsonString, _ := json.Marshal(ResponseMessage)
		fmt.Fprintf(w, "%s", jsonString)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = networkmanager.AddNetwork(n.Ssid, n.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ResponseMessage.Success = false
		ResponseMessage.Message = "Failed to add" + n.Ssid
		jsonString, _ := json.Marshal(ResponseMessage)
		fmt.Fprintf(w, "%s", jsonString)
		return
	}

	ResponseMessage.Success = true
	ResponseMessage.Message = "Successfully added" + n.Ssid

	jsonString, err := json.Marshal(ResponseMessage)

	fmt.Fprintf(w, "%s", jsonString)

}

// func CreatePersonalNetwork(w http.ResponseWriter, req *http.Request) {
// 	type network struct {
// 		Ssid     string `json:"ssid"`
// 		Password string `json:"password"`
// 	}
// 	var n network
// 	decoder := json.NewDecoder(req.Body)
// 	err := decoder.Decode(&n)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		ResponseMessage.Success = false
// 		ResponseMessage.Message = "Failed to add" + n.Ssid
// 		jsonString, _ := json.Marshal(ResponseMessage)
// 		fmt.Fprintf(w, "%s", jsonString)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	err = networkmanager.CreateAccessPoint(n.Ssid, n.Password)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		ResponseMessage.Success = false
// 		ResponseMessage.Message = "Failed to add" + n.Ssid
// 		jsonString, _ := json.Marshal(ResponseMessage)
// 		fmt.Fprintf(w, "%s", jsonString)
// 		return
// 	}
// 	ResponseMessage.Success = true
// 	ResponseMessage.Message = "Successfully added" + n.Ssid
// 	jsonString, err := json.Marshal(ResponseMessage)
// 	fmt.Fprintf(w, "%s", jsonString)
// }

func internetSharing() {
	wiredDevices, _ := networkmanager.GetDeviceByType(1)

	fmt.Println(wiredDevices)

	// Get first wired device (eth0)

	wiredDevice := wiredDevices[0].(map[string]interface{})["DevicePath"].(dbus.ObjectPath)

	networkmanager.InternetSharingOverEthernet(wiredDevice)
}

func scanWirelessNetwork(w http.ResponseWriter, req *http.Request) {
	wirelessDevices, _ := networkmanager.GetDeviceByType(2)
	// Get first wireless device (wlan0)
	wirelessDevice := wirelessDevices[0].(map[string]interface{})["DevicePath"].(dbus.ObjectPath)
	wirelessNetworks, err := networkmanager.GetAccessPoints(wirelessDevice)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ResponseMessage.Success = false
		ResponseMessage.Message = "Failed to create access point"
		jsonString, _ := json.Marshal(ResponseMessage)
		fmt.Fprintf(w, "%s", jsonString)
		return
	}
	jsonString, err := json.Marshal(wirelessNetworks)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", jsonString)
}

func createPersonalHotspot(w http.ResponseWriter, req *http.Request) {

	var n networkmanager.APNetwork

	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&n)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ResponseMessage.Success = false
		ResponseMessage.Message = "Failed to create access point"
		jsonString, _ := json.Marshal(ResponseMessage)
		fmt.Fprintf(w, "%s", jsonString)
		return
	}

	wirelessDevices, _ := networkmanager.GetDeviceByType(2)
	wirelessDevice := wirelessDevices[0].(map[string]interface{})["DevicePath"].(dbus.ObjectPath)

	wirelessNetwork := networkmanager.APNetwork{
		Ssid:     n.Ssid,
		Password: n.Password,
		Device:   wirelessDevice,
	}

	networkmanager.CreateAccessPoint(wirelessNetwork)
	w.Header().Set("Content-Type", "application/json")
	ResponseMessage.Success = true
	ResponseMessage.Message = "Successfully created access point"
	jsonString, err := json.Marshal(ResponseMessage)
	fmt.Fprintf(w, "%s", jsonString)
}

func main() {

	// internetSharing()

	//////////////////////////////

	// Get the list of wireless devices
	// devicess, _ := networkmanager.GetDeviceByType(2)

	// // get the first device's object path
	// device := devicess[0].(map[string]interface{})["DevicePath"].(dbus.ObjectPath)

	// settings, err := networkmanager.GetConnectionSettings("/org/freedesktop/NetworkManager/Settings/71")

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("SETTINGS", settings)

	// aps, err := networkmanager.GetAccessPoints("/org/freedesktop/NetworkManager/Devices/2")

	// if err != nil {
	// 	panic(err)
	// }

	// for _, attr := range wireless {
	// 	println(attr.(map[string]interface{})["State"].(uint32))
	// }

	// fmt.Println("WIRED", wired)

	http.HandleFunc("/settings", settings)

	// Connectivity
	http.HandleFunc("/connectivity", getConnectivity)

	// Device-related
	http.HandleFunc("/devices", getAllDevices)
	http.HandleFunc("/devices/wired", getWiredDevices)
	http.HandleFunc("/devices/wireless", getWirelessDevices)

	// Network-related
	http.HandleFunc("/add", addNetwork)
	http.HandleFunc("/create", createPersonalHotspot)
	http.HandleFunc("/scan", scanWirelessNetwork)

	// Start the server on port 8888

	http.ListenAndServe(":8888", nil)

}
