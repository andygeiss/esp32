# The ESP-IDF project, for the board itself. internal/device/controller.go is
# the program; `make firmware` transpiles it into main/controller.ino and this
# build compiles that against the Arduino core.
cmake_minimum_required(VERSION 3.16)

include($ENV{IDF_PATH}/tools/cmake/project.cmake)
project(esp32_controller)
