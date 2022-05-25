package networkmanager

import (
	"github.com/godbus/dbus/v5"
)

func GetWirelessDevices() ([]dbus.ObjectPath, error) {

}

func ListDevices() (interface{}, error) {
	obj := c.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager")
	devices, err := obj.GetProperty("org.freedesktop.NetworkManager.Devices")

	if err != nil {
		return nil, err
	}

	return devices.Value().([]dbus.ObjectPath), nil
}
