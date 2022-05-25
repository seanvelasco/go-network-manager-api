package main

import (
	"encoding/json"
	"fmt"
	"go-nm/networkmanager"
	"net/http"
)

// func settings(w http.ResponseWriter, req *http.Request) {
// 	savedconnections, err := networkmanager.ListSavedConnections()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Fprintf(w, "%s", "savedconnections")
// }

func connectivity(w http.ResponseWriter, req *http.Request) {
	connectivity, err := networkmanager.CheckConnectivity()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%d", connectivity)
}

func devices(w http.ResponseWriter, req *http.Request) {
	devices, err := networkmanager.ListDevices()
	if err != nil {
		panic(err)
	}
	jsonString, err := json.Marshal(devices)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", jsonString)
}

func wiredDevices(w http.ResponseWriter, req *http.Request) {
	devices, err := networkmanager.GetDeviceByType(1)

	if err != nil {
		fmt.Println(err)
	}
	jsonString, err := json.Marshal(devices)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", jsonString)
}

func wirelessDevices(w http.ResponseWriter, req *http.Request) {
	devices, err := networkmanager.GetDeviceByType(2)

	if err != nil {
		fmt.Println(err)
	}
	jsonString, err := json.Marshal(devices)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", jsonString)
}

var ResponseMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
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

func main() {

	settings, err := networkmanager.GetConnectionSettings("/org/freedesktop/NetworkManager/Settings/71")

	if err != nil {
		panic(err)
	}
	fmt.Println("SETTINGS", settings)

	// aps, err := networkmanager.GetAccessPoints("/org/freedesktop/NetworkManager/Devices/2")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("ACCESS POINTS", aps)

	// for _, attr := range wireless {
	// 	println(attr.(map[string]interface{})["State"].(uint32))
	// }

	// fmt.Println("WIRED", wired)

	// http.HandleFunc("/settings", settings)
	http.HandleFunc("/connectivity", connectivity)
	http.HandleFunc("/devices", devices)
	http.HandleFunc("/devices/wired", wiredDevices)
	http.HandleFunc("/devices/wireless", wirelessDevices)

	http.HandleFunc("/add", addNetwork)

	http.ListenAndServe(":8888", nil)

}
