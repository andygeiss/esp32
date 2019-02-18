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

First download and install the latest [Arduino IDE](https://www.arduino.cc/en/Main/Software) into <code>/opt/arduino</code> or change <code>INO_IDE_PATH</code> in the <code>Makefile</code>
and necessary packages.
 
- **Ubuntu**:

        sudo apt-get install -y bison flex git gperf libncurses-dev make python
	
- **Manjaro**:

        sudo pip install pyserial

Next run the ESP32 SDK-Installation:

    make packages

Look at the [examples](https://github.com/andygeiss/esp32/tree/master/examples) for more information.

## Develop, Test and Build

Change the Arduino port to your current settings by changing the <code>Makefile</code>:

    INO_PORT="/dev/ttyUSB0"

Also set the <code>SSID</code> and <code>PASS</code> strings at <code>application/device/controller.go</code> to your WiFi Access Point and start the build process by using:

    make
    
Run the binary at <code>build/device-${ARCH}</code> to simulate your ESP32 device locally.

    Connecting to WiFi ...... Connected!

## Deploy

Finally use the following command to deploy the encrypted <code>device.img</code> to your real ESP32 device.

    make flash

**Important**: PLease ensure that the current user is in the <code>dialout</code> group. Or you will receive a <code>permission denied</code>.

This will create the following output:

    2018/08/05 16:04:22 Flashing ...
    esptool.py v2.3.1
    Connecting....
    Chip is ESP32D0WDQ6 (revision 1)
    Features: WiFi, BT, Dual Core
    Uploading stub...
    Running stub...
    Stub running...
    Changing baud rate to 921600
    Changed.
    Configuring flash size...
    Compressed 607456 bytes to 366770...
    Wrote 607456 bytes (366770 compressed) at 0x00010000 in 5.9 seconds (effective 816.9 kbit/s)...
    Hash of data verified.
    
    Leaving...
    Hard resetting via RTS pin...
    2018/08/05 16:04:22 Done.
