package networkmanager

import "github.com/godbus/dbus/v5"

func GetConnectionSettings(path dbus.ObjectPath) (interface{}, error) {

	obj := c.Object("org.freedesktop.NetworkManager", path)

	var setting map[string]map[string]dbus.Variant
	obj.Call("org.freedesktop.NetworkManager.Settings.Connection.GetSettings", 0).Store(&setting)

	rv := make(ConnectionSettings)

	for k1, v1 := range setting {
		rv[k1] = make(map[string]interface{})

		for k2, v2 := range v1 {
			rv[k1][k2] = v2.Value()
		}
	}

	return rv, nil
}

func activateConnection(connection interface{}, path dbus.ObjectPath) (interface{}, error) {

	obj := c.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager")

	var activeConnection []interface{}
	obj.Call("org.freedesktop.NetworkManager.ActivateConnection", 0, connection, path, "/").Store(&activeConnection)

	return activeConnection, nil
}
