APPNAME=$(shell basename `pwd`)
ARCH=$(shell uname -m)
LDFLAGS="-s"

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
	@go build -ldflags $(LDFLAGS) -o build/$(APPNAME)-$(ARCH) main/main.go
	@esp32-transpiler -source $(INO_SOURCE_FILE) -target $(INO_SKETCH_FILE)
	@$(INO_TOOLS_BUILD) --ide_path=$(INO_IDE_PATH) -d $(INO_HARDWARE_PATH) -b $(INO_BOARD) -w all -o $(INO_SKETCH_IMAGE) $(INO_SKETCH_FILE)
	
build: build/$(APPNAME)

clean:
	@rm -f build/*
	
flash:
	@$(INO_TOOLS_FLASH) --chip $(INO_BOARD) --port $(INO_PORT) --baud $(INO_BAUD) write_flash -fm dio -fs 4MB -ff 40m 0x00010000 $(INO_SKETCH_IMAGE)
	
packages:
	@go get -u github.com/andygeiss/esp32-controller
	@go get -u github.com/andygeiss/esp32-transpiler
	@go install -u github.com/andygeiss/esp32-transpiler
	@rm -rf $(INO_SDK_PATH)
	@mkdir -p $(INO_SDK_PATH)
	@git clone --recursive https://github.com/espressif/arduino-esp32.git $(INO_SDK_PATH)
	@curl -SSL https://dl.espressif.com/dl/xtensa-esp32-elf-linux64-1.22.0-73-ge28a011-5.2.0.tar.gz > $(INO_SDK_PATH)/tools/xtensa.tar.gz
	@tar xzf $(INO_SDK_PATH)/tools/xtensa.tar.gz -C $(INO_SDK_PATH)/tools/
	@rm -f $(INO_SDK_PATH)/tools/xtensa.tar.gz
	@git clone --recursive https://github.com/espressif/esp-idf.git $(INO_SDK_PATH)/framework
	@mkdir -p $(INO_SDK_PATH)/tools/esptool
	@ln -sf $(INO_SDK_PATH)/tools/esptool.py $(INO_SDK_PATH)/tools/esptool/
	
run:
	@./build/$(APPNAME)-$(ARCH)
	
test:
	@go test -v ./...
