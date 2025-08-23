all: setup bios.out

setup: clean

BIOS_DIR=backend/bios
BUILD_DIR=build
LIB_DIR=backend/lib

setup:
	mkdir build

crt0.o:
	ca65 $(BIOS_DIR)/crt0.s -o $(BUILD_DIR)/crt0.o

base.lib: crt0.o
	cp $(LIB_DIR)/base.lib $(BUILD_DIR)/base.lib # take a copy as it gets replaced
	ar65 a $(BUILD_DIR)/base.lib $(BUILD_DIR)/crt0.o

bios.s:
	/Users/carl/dev/cc65/bin/cc65 $(BIOS_DIR)/bios.c -o $(BUILD_DIR)/bios.s

bios.o: bios.s
	/Users/carl/dev/cc65/bin/ca65 -o $(BUILD_DIR)/bios.o $(BUILD_DIR)/bios.s

bios.out: bios.o base.lib
	/Users/carl/dev/cc65/bin/ld65 -C $(BIOS_DIR)/system.cfg $(BUILD_DIR)/bios.o $(BUILD_DIR)/base.lib -o $(BUILD_DIR)/bios.out

clean:
	rm -rf build