package wifi_test

import (
	"testing"

	"github.com/andygeiss/goat/business/controller/wifi"
)

func TestWifiBegin(t *testing.T) {
	ssid := "test"
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.Begin(ssid)
	if wifi.CurrentStatus != wifi.StatusConnected {
		t.Error("Status should be connected!")
	}
}

func TestWifiBeginEncrypted(t *testing.T) {
	ssid := "test"
	passphrase := "passphrase"
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.BeginEncrypted(ssid, passphrase)
	if wifi.CurrentStatus != wifi.StatusConnected {
		t.Error("Status should be connected!")
	}
}
func TestWifiDisBegin(t *testing.T) {
	ssid := "test"
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.Begin(ssid)
	wifi.Disconnect()
	if wifi.CurrentStatus != wifi.StatusIdle {
		t.Error("Status should be idle!")
	}
}

func TestWifiRSSI(t *testing.T) {
	ssid := "test"
	wifi.CurrentRSSI = -1
	wifi.Begin(ssid)
	if wifi.RSSI() == -1 {
		t.Error("Signal strength should be greater than -1!")
	}
}
func TestWifiSSID(t *testing.T) {
	ssid := "test"
	wifi.CurrentSSID = ""
	wifi.Begin(ssid)
	if wifi.SSID() == "" {
		t.Error("SSID should not be empty!")
	}
}
