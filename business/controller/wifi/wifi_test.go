package wifi_test

import (
	"testing"

	. "github.com/andygeiss/assert"

	"github.com/andygeiss/esp32/business/controller/wifi"
)

func TestWifiBegin(t *testing.T) {
	ssid := "test"
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.Begin(ssid)
	Assert(t, wifi.CurrentStatus, IsEqual(wifi.StatusConnected))
}

func TestWifiBeginEncrypted(t *testing.T) {
	ssid := "test"
	passphrase := "passphrase"
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.BeginEncrypted(ssid, passphrase)
	Assert(t, wifi.CurrentStatus, IsEqual(wifi.StatusConnected))
}
func TestWifiDisBegin(t *testing.T) {
	ssid := "test"
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.Begin(ssid)  // StatusConnected
	wifi.Disconnect() // back to idle?
	Assert(t, wifi.CurrentStatus, IsEqual(wifi.StatusIdle))
}

func TestWifiRSSIShouldBeNotMinusOne(t *testing.T) {
	ssid := "test"
	wifi.CurrentRSSI = -1
	wifi.Begin(ssid)
	Assert(t, wifi.RSSI(), IsNotEqual(-1))
}
func TestWifiSSIDShouldNotBeEmpty(t *testing.T) {
	ssid := "test"
	wifi.CurrentSSID = ""
	wifi.Begin(ssid)
	Assert(t, wifi.SSID(), IsNotEqual(""))
}
