# Binary settings
GITUSER=andygeiss
APPNAME=$(shell cat APPNAME)
BUILD=$(shell date -u +%Y%m%d%H%M%S)
VERSION=$(shell cat VERSION)
LDFLAGS="-s -X main.APPNAME=$(APPNAME) -X main.BUILD=$(BUILD) -X main.VERSION=$(VERSION)"
TS=$(shell date -u '+%Y/%m/%d %H:%M:%S')

INO_BAUD="921600"
INO_BOARD="esp32"
INO_HARDWARE_PATH="/home/$(USER)/Arduino/hardware"
INO_IDE_PATH="/opt/arduino"
INO_MANUFACTURER=espressif
INO_MAPPING="infrastructure/ino/mapping.json"
INO_PORT="/dev/ttyUSB0"
INO_SDK_PATH="$(INO_HARDWARE_PATH)/$(INO_MANUFACTURER)/$(INO_BOARD)"
INO_SKETCH_FILE=build/device.ino
INO_SKETCH_IMAGE=build/device.img
INO_SOURCE_FILE=application/device/controller.go
INO_TOOLS_PATH="$(INO_HARDWARE_PATH)/espressif/$(INO_BOARD)/tools"
INO_TOOLS_BUILD="$(INO_TOOLS_PATH)/build.py"
INO_TOOLS_FLASH="$(INO_TOOLS_PATH)/esptool.py"

TRANSPILER_BINARY_FILE=build/transpile
TRANSPILER_SOURCE_FILE=platform/transpile/main.go
TRANSPILER_MAPPING=infrastructure/ino/mapping.json

X86_SOURCE_FILE=platform/device/main.go
X86_BINARY_FILE=build/device-x86_64.bin

all: clean test build

build/$(APPNAME):
	@echo $(TS) Building $(TRANSPILER_BINARY_FILE) ...
	@go build -ldflags $(LDFLAGS) -o $(TRANSPILER_BINARY_FILE) $(TRANSPILER_SOURCE_FILE)
	@echo $(TS) Building $(X86_BINARY_FILE) ...
	@go build -ldflags $(LDFLAGS) -o $(X86_BINARY_FILE) $(X86_SOURCE_FILE)
	@echo $(TS) Building $(INO_SKETCH_FILE) ...
	@./$(TRANSPILER_BINARY_FILE) -source $(INO_SOURCE_FILE) -mapping $(TRANSPILER_MAPPING) -target $(INO_SKETCH_FILE)
	@echo $(TS) Building $(INO_SKETCH_IMAGE) ...
	@$(INO_TOOLS_BUILD) --ide_path=$(INO_IDE_PATH) -d $(INO_HARDWARE_PATH) -b $(INO_BOARD) -w all -o $(INO_SKETCH_IMAGE) $(INO_SKETCH_FILE)
	@echo $(TS) Done.

build: build/$(APPNAME)

clean:
	@echo $(TS) Cleaning up previous build ...
	@rm -f build/*
	@echo $(TS) Done.

flash:
	@echo $(TS) Flashing ...
	@$(INO_TOOLS_FLASH) --chip $(INO_BOARD) --port $(INO_PORT) --baud $(INO_BAUD) write_flash -fm dio -fs 4MB -ff 40m 0x00010000 $(INO_SKETCH_IMAGE)
	@echo $(TS) Done.

init:
	@echo $(TS) Creating initial commit ...
	@rm -rf .git
	@git init
	@git add .
	@git commit -m "Initial commit"
	@git remote add origin git@github.com:$(GITUSER)/$(APPNAME).git
	@git push -u --force origin master
	@echo $(TS) Done.

install:
	@echo $(TS) Installing SDK ...
	@rm -rf $(INO_SDK_PATH)
	@mkdir -p $(INO_SDK_PATH)
	@git clone --recursive https://github.com/espressif/arduino-esp32.git $(INO_SDK_PATH)
	@curl -SSL https://dl.espressif.com/dl/xtensa-esp32-elf-linux64-1.22.0-73-ge28a011-5.2.0.tar.gz > $(INO_SDK_PATH)/tools/xtensa.tar.gz
	@tar xzf $(INO_SDK_PATH)/tools/xtensa.tar.gz -C $(INO_SDK_PATH)/tools/
	@rm -f $(INO_SDK_PATH)/tools/xtensa.tar.gz
	@git clone --recursive https://github.com/espressif/esp-idf.git $(INO_SDK_PATH)/framework
	@sudo apt-get update && sudo apt-get install -y bison flex git gperf libncurses-dev make python python-serial
	@echo $(TS) Done.

test:
	@echo $(TS) Testing ...
	@go test -v ./...
	@echo $(TS) Done.
