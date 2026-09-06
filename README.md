# ESP32

[![License](https://img.shields.io/github/license/andygeiss/esp32)](https://github.com/andygeiss/esp32/blob/master/LICENSE)
[![Releases](https://img.shields.io/github/v/release/andygeiss/esp32)](https://github.com/andygeiss/esp32/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/esp32)](https://goreportcard.com/report/github.com/andygeiss/esp32)
[![Maintainability](https://api.codeclimate.com/v1/badges/25eafcae9797d5e8a3de/maintainability)](https://codeclimate.com/github/andygeiss/esp32/maintainability)

Write the code for your ESP32 in Go, run it on your own machine, then flash it to
the device. This is for anyone who has grown tired of the Arduino IDE's only way to
find a bug: compile, flash, and stare at the serial monitor.

## Run it

    git clone https://github.com/andygeiss/esp32.git
    cd esp32
    make run

That builds the controller in `internal/device/controller.go` and runs it here,
the way the ESP32 would:

    Connecting to WiFi .... Connected!
    Connecting to [HOSTNAME:3000)] ... Connected!
    GET /index.html HTTP/1.0

Press Ctrl-C to stop. A second Ctrl-C kills the process, which is the way out
when a loop is in the middle of a long `timer.Delay`.

The binary takes two flags:

    esp32 -n 5        # stop after five loops; 0, the default, runs until you stop it
    esp32 -version    # print the version and exit

The controller's output goes to stdout, so `esp32 > serial.log` captures exactly
what the serial monitor would show. Everything else — errors, usage — goes to
stderr.

You can install it too:

    go install github.com/andygeiss/esp32@latest

That gives you the controller *from this repository*. Yours lives in
`internal/device/controller.go`, so the real path is clone, edit, `make`.

## Write your own controller

Edit `internal/device/controller.go`. `Setup` runs once, `Loop` runs forever —
the same two functions an Arduino sketch has, and here they are plain
package-level functions. Set `ssid` and `pass` to your own access point while
you are in there.

`examples/blink/controller.go` is a smaller starting point that flashes the LED
on pin 2. Copy it over `internal/device/controller.go` and change its package
line to `device`. It holds nothing else, so that is the whole edit.

**The transpiler reads that one file, and it understands only a small subset of
[the Go language spec](https://go.dev/ref/spec).** Stick to plain functions,
constants, package-level variables and calls into `esp32-controller`. **A type
declaration and a method are both outside it** — that is why `Setup` and `Loop`
are functions rather than methods on a `Controller`, and why the interface the
simulator runs lives next door in `internal/device/program.go`, which the
transpiler never reads. The
[mapping](https://github.com/andygeiss/esp32-transpiler/blob/master/internal/transpile/handlers/mapping.go)
and the
[tests](https://github.com/andygeiss/esp32-transpiler/blob/master/internal/transpile/service_test.go)
say what it can do today. Anything else — the standard library, generics, a
context — belongs in another package.

Get it wrong and `make firmware` stops with the construct and the line it sits
on, and writes no sketch:

    esp32-transpiler: internal/device/controller.go:19:1: a type declaration is not supported

That refusal is the point. A construct the tool passed over in silence would
reach the board as a sketch that compiles and does something other than the Go
said, which is exactly what `go test` was run to rule out.

## Test it

    make          # every gate: format, vet, staticcheck, govulncheck, race tests, build
    make test     # just the tests, while you work

There is no CI server. `make check` is the gate against your working tree and
`make ci` runs the same gates against the commit, so run that before you push.

Tests wrap the controller in a [`synctest`](https://pkg.go.dev/testing/synctest)
bubble. A loop that waits five seconds on the device takes no time at all here.

## Install ESP-IDF

You only need this to build and flash a real image. Running and testing the
controller on your own machine needs nothing but Go.

ESP-IDF lives outside the tree, in the same place
[esp32-watch](https://github.com/andygeiss/esp32-watch) keeps it, so one install
serves both projects:

    git clone --depth 1 --shallow-submodules --recursive -b v5.5 \
        https://github.com/espressif/esp-idf.git ~/esp/esp-idf
    ~/esp/esp-idf/install.sh esp32

Then export it into the shell you build in. That is not a permanent change to
your PATH, so every new terminal needs it again:

    . ~/esp/esp-idf/export.sh

That is the whole install. The Arduino core arrives on its own: it is a pinned
ESP-IDF component in `firmware/main/idf_component.yml`, and the component
manager downloads it the first time you build.

## Build and flash

    make firmware    # transpile the controller, then build the image
    make flash       # write it to the board, then show the serial log

`make firmware` turns `internal/device/controller.go` into
`firmware/main/controller.ino` and hands that to ESP-IDF. Ctrl-] leaves the
monitor `make flash` opens.

Both targets take two overrides:

    make firmware TARGET=esp32s3          # build for a different chip
    make flash ESPPORT=/dev/ttyUSB0       # name the serial port yourself

`idf.py` finds the board on its own most of the time, so `ESPPORT` is for when
it guesses wrong. On Linux **your user has to be in the `dialout` group**, or
flashing stops at `permission denied`.

The image is about 900 KB, nearly all of it the Arduino core and the WiFi stack.
`firmware/partitions.csv` gives the app 2 MB of the board's 4 MB so there is room
to write a real controller; the built-in table would leave 12% free before you
started.

## Why there are no environment variables

The SSID, the passphrase and the host are constants in
`internal/device/controller.go`, not settings. The transpiler turns them into
`const char*` in the sketch, so they are baked into the image before the device
ever boots. There is nothing left to configure at run time.

## What is in here

    esp32/
    ├── examples/
    │   └── blink/
    │       ├── controller.go        ← a starter controller you can copy over
    │       └── controller_test.go
    ├── firmware/
    │   ├── CMakeLists.txt           ← the ESP-IDF project
    │   ├── dependencies.lock        ← pins the Arduino core; checked in
    │   ├── main/
    │   │   ├── CMakeLists.txt
    │   │   ├── idf_component.yml    ← the Arduino core, pinned
    │   │   └── sketch.cpp           ← includes the transpiled controller.ino
    │   ├── partitions.csv
    │   └── sdkconfig.defaults
    ├── go.mod
    ├── go.sum
    ├── internal/
    │   ├── device/
    │   │   ├── controller.go        ← the file you edit; the only file transpiled
    │   │   ├── controller_test.go
    │   │   └── program.go           ← wraps it in the interface the simulator runs
    │   └── simulator/
    │       ├── simulator.go         ← setup once, then loop, on this machine
    │       └── simulator_test.go
    ├── LICENSE
    ├── main.go
    ├── main_test.go
    ├── Makefile
    ├── README.md
    └── SPEC.md

Three repositories make this work. This one holds your controller and the
toolchain around it.
[esp32-controller](https://github.com/andygeiss/esp32-controller) is the Arduino
API written in Go, and
[esp32-transpiler](https://github.com/andygeiss/esp32-transpiler) rewrites your
Go into the C++ the device compiles. `make firmware` fetches the transpiler with
`go run`, so there is nothing to install. The C++ it writes is Arduino, which
ESP-IDF compiles through the `espressif/arduino-esp32` component.

## Baseline deviations

This project follows the [engineering baseline](https://github.com/andygeiss/baseline)
as a **CLI tool**, the closest of its three project types. Below is every rule it
waives, every rule it meets by another route, and the one dependency it has to
justify.

- **A CLI terminates** (`project-types/cli-tool.md`) — waived 2026-09-06 by Andy.
  `esp32` mirrors the device: set up once, then loop until something stops it. The
  baseline has no project type for firmware, and a tool that ended after one pass
  would not show you what the device does. Contained by `-n`: every test and every
  gate runs a bounded number of loops, and Ctrl-C ends an unbounded one.

- **Everything under `internal/`** (`patterns/go-project-layout.md` rule 2) —
  waived 2026-09-06 by Andy. `examples/blink` exists to be opened and copied, so
  it sits where a reader looks first. Contained to `examples/`: nothing imports
  it, and `make check` covers it like any other package.

- **The Makefile is the baseline's** (`stack/makefile.md`) — waived 2026-09-06 by
  Andy, on two counts. Rule 3 allows the `firmware` and `flash` targets, but not
  the `TARGET` and `TRANSPILER` variables above them; those are what let you
  build for another chip without editing the file. And rule 4 floats dev tools at
  `@latest`, while `TRANSPILER` names an exact version: the transpiler's output
  is the firmware, not a pass or a fail, so a floating version would change what
  gets flashed without a commit.

- **Ports and adapters** (`patterns/go-ports-adapters.md`) — met by another route.
  `esp32-controller` is already the hand-written fake this pattern asks for: on
  the device its calls become Arduino API calls, and here they move numbers around
  in a map. There is no second implementation left to hide behind a port.

- **`github.com/andygeiss/esp32-controller`** is the only Go dependency, and it is
  not on the baseline's approved list. It is the Arduino API spelled in Go, and
  the transpiler rewrites each call into its C++ equivalent. Nothing in the
  standard library reaches a GPIO pin. `espressif/arduino-esp32` in
  `firmware/main/idf_component.yml` is the other half of that same story — the
  C++ those calls land in — and it is an ESP-IDF component, which the baseline's
  Go approved list does not speak to either way.
