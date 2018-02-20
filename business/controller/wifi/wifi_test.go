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
	ipv4 := wifi.NewIPv4Address(127, 0, 0, 1)
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.BeginEncrypted(ssid, passphrase)
	Assert(t, wifi.CurrentStatus, IsEqual(wifi.StatusConnected))
	Assert(t, wifi.CurrentLocalIP, IsEqual(ipv4))
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
