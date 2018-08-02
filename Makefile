APPNAME=$(shell basename `pwd`)
ARCH=$(shell uname -m)
LDFLAGS="-s"
TS=$(shell date -u '+%Y/%m/%d %H:%M:%S')

INO_BAUD="921600"
INO_BOARD="esp32"
INO_HARDWARE_PATH="/home/$(USER)/Arduino/hardware"
INO_IDE_PATH="/opt/arduino"
INO_MANUFACTURER=espressif
INO_PORT="/dev/ttyUSB0"
INO_SDK_PATH="$(INO_HARDWARE_PATH)/$(INO_MANUFACTURER)/$(INO_BOARD)"
INO_SKETCH_FILE=build/device.ino
INO_SKETCH_IMAGE=build/device.img
INO_SOURCE_FILE=device/controller.go
INO_TOOLS_PATH="$(INO_HARDWARE_PATH)/espressif/$(INO_BOARD)/tools"
INO_TOOLS_BUILD="$(INO_TOOLS_PATH)/build.py"
INO_TOOLS_FLASH="$(INO_TOOLS_PATH)/esptool.py"

all: clean test build

build/$(APPNAME):
	@echo $(TS) Building local device ...
	@go build -ldflags $(LDFLAGS) -o build/$(APPNAME)-$(ARCH) main/main.go
	@echo $(TS) Building $(INO_SKETCH_FILE) ...
	@esp32-transpiler -source $(INO_SOURCE_FILE) -target $(INO_SKETCH_FILE)
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

packages:
	@echo $(TS) Installing Go packages ...
	@go get -u github.com/andygeiss/esp32-controller
	@go get -u github.com/andygeiss/esp32-transpiler
	@echo $(TS) Installing SDK ...
	@rm -rf $(INO_SDK_PATH)
	@mkdir -p $(INO_SDK_PATH)
	@git clone --recursive https://github.com/espressif/arduino-esp32.git $(INO_SDK_PATH)
	@curl -SSL https://dl.espressif.com/dl/xtensa-esp32-elf-linux64-1.22.0-73-ge28a011-5.2.0.tar.gz > $(INO_SDK_PATH)/tools/xtensa.tar.gz
	@tar xzf $(INO_SDK_PATH)/tools/xtensa.tar.gz -C $(INO_SDK_PATH)/tools/
	@rm -f $(INO_SDK_PATH)/tools/xtensa.tar.gz
	@git clone --recursive https://github.com/espressif/esp-idf.git $(INO_SDK_PATH)/framework
	@sudo apt-get update && sudo apt-get install -y bison flex git gperf libncurses-dev make python python-serial
	@echo $(TS) Adding symbolic link ...
	@mkdir -o $(INO_SDK_PATH)/tools/esptool
	@ln -sf $(INO_SDK_PATH)/tools/esptool.py $(INO_SDK_PATH)/tools/esptool/
	@echo $(TS) Done.

run:
	@echo $(TS) Running $(APPNAME) ...
	@./build/$(APPNAME)-$(ARCH)
	@echo $(TS) Done.

test:
	@echo $(TS) Testing ...
	@go test -v ./...
	@echo $(TS) Done.
