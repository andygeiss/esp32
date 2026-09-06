// The transpiled controller, compiled as part of this component.
//
// There is no app_main below. CONFIG_AUTOSTART_ARDUINO puts the Arduino core's
// own entry point in the build, and that is what calls setup() once and then
// loop() over and over.

// Arduino.h comes first: the transpiler emits an include only for the wifi
// package, so a controller that just blinks an LED produces a sketch with no
// includes at all, and Serial, delay and pinMode would be undeclared.
#include <Arduino.h>

#if !__has_include("controller.ino")
#error "no controller.ino - run `make firmware`, which transpiles internal/device/controller.go into it"
#endif

#include "controller.ino"
