ROOT_DIR=$(shell pwd)
MODULE_DIR=$(ROOT_DIR)/cmd
BUILD_DIR=$(ROOT_DIR)/bin
EXE=$(BUILD_DIR)/ccwc

build: clean
	@ mkdir $(BUILD_DIR)
	go build -o $(EXE)

clean:
	rm -rf $(BUILD_DIR)

test:
	@ cd $(MODULE_DIR); go test -v; cd $(ROOT_DIR)


