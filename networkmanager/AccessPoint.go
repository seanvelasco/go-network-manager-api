package networkmanager

import (
	"github.com/godbus/dbus/v5"
)

func (a *accessPoint) GetSSID() (string, error) {

	ssid, err := a.getSliceByteProperty("org.freedesktop.NetworkManager.AccessPoint.Ssid")

	if err != nil {
		return "", err
	}

	return string(ssid), nil
}

func (a *accessPoint) GetFrequency() (uint32, error) {

	freq, err := a.obj.GetProperty("org.freedesktop.NetworkManager.AccessPoint.Frequency")

	if err != nil {
		return 0, err
	}

	return freq.Value().(uint32), nil
}

func (a *accessPoint) GetStrength() (uint8, error) {

	strength, err := a.obj.GetProperty("org.freedesktop.NetworkManager.AccessPoint.Strength")

	if err != nil {
		return 0, err
	}

	return strength.Value().(uint8), nil
}

func (a *accessPoint) GetFlags() (uint32, error) {

	flags, err := a.getUint32Property("org.freedesktop.NetworkManager.AccessPoint.Flags")

	if err != nil {
		return 0, err
	}

	return uint32(flags), nil
}

type Nm80211Mode uint32

func (a *accessPoint) GetMode() (uint32, error) {

	mode, err := a.obj.GetProperty("org.freedesktop.NetworkManager.AccessPoint.Mode")

	if err != nil {
		return 0, err
	}

	return mode.Value().(uint32), nil
}

func (a *accessPoint) GetHWAddress() (string, error) {

	hwaddress, err := a.obj.GetProperty("org.freedesktop.NetworkManager.AccessPoint.HwAddress")

	if err != nil {
		return "", err
	}

	return hwaddress.Value().(string), nil

}

// MaxBitrate
func (a *accessPoint) GetMaxBitrate() (uint32, error) {

	maxBitrate, err := a.obj.GetProperty("org.freedesktop.NetworkManager.AccessPoint.MaxBitrate")
	if err != nil {
		return 0, err
	}
	return maxBitrate.Value().(uint32), nil

}

// scanWirelessNetworks
func GetAccessPointInfo(path dbus.ObjectPath) (interface{}, error) {

	obj := c.Object("org.freedesktop.NetworkManager", path)

	ap := &accessPoint{
		dbusBase: dbusBase{
			conn: c,
			obj:  obj,
		},
	}

	ssid, _ := ap.GetSSID()
	frequency, _ := ap.GetFrequency()
	strength, _ := ap.GetStrength()
	hwaddress, _ := ap.GetHWAddress()
	flags, _ := ap.GetFlags()
	maxbitrate, _ := ap.GetMaxBitrate()

	mode, err := ap.GetMode()

	if err != nil {
		return nil, err
	}

	// Join ssid and freq into one object
	return map[string]interface{}{
		"ssid":       ssid,
		"frequency":  frequency,
		"strength":   strength,
		"hwaddress":  hwaddress,
		"flags":      flags,
		"mode":       mode,
		"maxbitrate": maxbitrate,
	}, nil

}

func createAccessPoint(path dbus.ObjectPath) *accessPoint {

	obj := c.Object("org.freedesktop.NetworkManager", path)

	ap := &accessPoint{
		dbusBase: dbusBase{
			conn: c,
			obj:  obj,
		},
	}

	return ap
}
