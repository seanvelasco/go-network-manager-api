package networkmanager

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

func GetAccessPoints(path dbus.ObjectPath) (interface{}, error) {

	obj := c.Object("org.freedesktop.NetworkManager", path)

	aps, err := obj.GetProperty("org.freedesktop.NetworkManager.Device.Wireless.AccessPoints")

	if err != nil {
		return nil, err
	}

	var rv []interface{}

	for _, ap := range aps.Value().([]dbus.ObjectPath) {
		settings, err := GetAccessPointInfo(ap)
		if err != nil {
			fmt.Println(err)
		}
		rv = append(rv, settings)
	}

	if err != nil {
		return nil, err
	}

	return rv, nil

}
