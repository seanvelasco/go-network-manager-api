package networkmanager

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

func AddNetwork(ssid, password string) (interface{}, error) {

	settings := Settings{
		"connection": {
			"id":   ssid,
			"type": "802-11-wireless",
		},
		"802-11-wireless": {
			"ssid": []byte(ssid),
			"mode": "infrastructure",
		},
		"802-11-wireless-security": {
			"key-mgmt": "wpa-psk",
			"psk":      password,
		},
		"ipv4": {
			"method": "auto",
		},
		"ipv6": {
			"method": "auto",
		},
	}

	obj := c.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/Settings")

	view := obj.Call("org.freedesktop.NetworkManager.Settings.AddConnection", 0, settings) //.Store(&id)

	if view.Err != nil {
		return 0, view.Err
	}

	return view.Body[0], nil
}

func ForgetNetwork(path dbus.ObjectPath) (interface{}, error) {
	obj := c.Object("org.freedesktop.NetworkManager", path)
	state := obj.Call("org.freedesktop.NetworkManager.Settings.Connection.Delete", 0)

	if state.Err != nil {
		return nil, fmt.Errorf("Network does not exist")
	}

	return state.Body, nil
}
