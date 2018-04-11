# ESP32

[![](https://goreportcard.com/badge/github.com/andygeiss/esp32)](https://goreportcard.com/report/github.com/andygeiss/esp32)

Build your own toolchain to develop, test, build and finally deploy a Golang controller to your ESP32 device.

## Purpose

The [Arduino IDE](https://www.arduino.cc/en/Main/Software) is easy to use.
But I faced problems like maintainability and testability at more complicated IoT projects.
I needed to compile and flash the ESP32 before testing my code functionality by doing it 100% manually.

This solution transpiles Golang into Arduino code, which can be compiled to an image by using the ESP32 toolchain.
Now I am able to use a fully automated testing approach instead of doing it 100% manually.

## Process

    +--------+    +---------+    +----------+
    |  Test  +---->  Build  +---->  Deploy  |
    +--------+    +---------+    +----------+

              make                make flash

**Important**: The Transpiler only supports a small subset of the [Golang Language Specification](https://golang.org/ref/spec). 

## Installation

First download and install the latest [Arduino IDE](https://www.arduino.cc/en/Main/Software) into <code>/opt/arduino</code> or change <code>INO_IDE_PATH</code> in the <code>Makefile</code>.

Next run the ESP32 SDK-Installation:

    make packages

## Develop, Test and Build

Modify the Controller at <code>application/device/controller.go</code> and start the build process by using: 

    make

Run the binary at <code>build/device-${ARCH}</code> to simulate your ESP32 device locally.

    2018/03/14 15:05:09 Running esp32 ...
    Setting up PIN 2 -> OUTPUT
    Done.
      Write PIN 2 -> HIGH
      Write PIN 2 -> LOW
      Write PIN 2 -> HIGH
      Write PIN 2 -> LOW
      Write PIN 2 -> HIGH
      Write PIN 2 -> LOW

## Deploy

Finally use the following command to deploy the encrypted <code>device.img</code> to your real ESP32 device.

    make flash
