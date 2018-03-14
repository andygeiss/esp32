# ESP32

[![](https://goreportcard.com/badge/github.com/andygeiss/esp32)](https://goreportcard.com/report/github.com/andygeiss/esp32)

Build your own toolchain to develop, test, build and finally deploy a Golang controller to your ESP32 device.

## Why Go? Arduino is super easy!

The [Arduino IDE](https://www.arduino.cc/en/Main/Software) sure is easy to use.
But lets talk about maintainability at complicated projects like "Implementing a Robot with X sensors, WiFi and MQTT"?

I simply want to ensure that the functionality of my code is working BEFORE flashing the ESP32!

## Examples

Look at the following Arduino [ESP32 LED Blink Example Code](https://circuits4you.com/2018/02/02/esp32-led-blink-example/):

    #define LED 2

    void setup() {
        pinMode(LED,OUTPUT);
    }

    void loop() {
        delay(500);
        digitalWrite(LED,HIGH);
        delay(500);
        digitalWrite(LED,LOW);
    }

This could be implemented in Golang (see <code>application/device/controller</code>) like that:


    func (c *Controller) Setup() error {
        digital.PinMode(2, digital.ModeOutput)
        return nil
    }

    func (c *Controller) Loop() error {
        timer.Delay(500)
        digital.Write(2, digital.High)
        timer.Delay(500)
        digital.Write(2, digital.Low)
        return nil
    }

## Process

    +--------+    +---------+    +----------+
    |  Test  +---->  Build  +---->  Deploy  |
    +--------+    +---------+    +----------+

              make                make flash

**Important**: The Transpiler only supports a small subset of the [Golang Language Specification](https://golang.org/ref/spec). Look at the tests in <code>infrastructure/ino/worker_test.go</code> for more information.

## Installation

First download and install the latest [Arduino IDE](https://www.arduino.cc/en/Main/Software) into <code>/opt/arduino</code> or change <code>INO_IDE_PATH</code> in the <code>Makefile</code>. Next run the ESP32 SDK-Installation:

    make install

## Develop, Test and Build

Modify the Controller at <code>application/device/controller.go</code> and start the build process by using: 

    make

This will create the following binary outputs:

    build/
    ├── device.img
    ├── device.ino
    ├── device-x86_64.bin
    └── transpile

Run <code>device-x86_64.bin</code> to simulate your ESP32 device locally. 

## Deploy

Finally use the following command to deploy the encrypted <code>device.img</code> to your real ESP32 device.

    make flash
