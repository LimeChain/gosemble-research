TINYGO_DIR = ../tinygo
CURRENT_DIR = $(shell pwd)

# build with the system installed TinyGo 
all:
	@tinygo build \
	-wasm-abi=generic \
	-scheduler=none \
	-gc=none \
	-o=build/dev_runtime.wasm \
	runtime/runtime.go
	@echo "Standard TinyGo build, done!"

# build with the TinyGo fork
.PHONY: build
build:
	@docker build --tag polkawasm/tinygo:0.25.0 -f $(TINYGO_DIR)/Dockerfile.polkawasm $(TINYGO_DIR)
	@docker run --rm -v $(CURRENT_DIR):/src/examples/wasm/gosemble -w /src/examples/wasm/gosemble polkawasm/tinygo:0.25.0 /bin/bash \
	-c "tinygo build -target=polkawasm -o /src/examples/wasm/gosemble/build/dev_runtime.wasm /src/examples/wasm/gosemble/runtime/runtime.go"
# wat2wasm -o build/dev_runtime.wasm build/dev_runtime.wat 

run:
	go run main.go

test:
	go test -v runtime/runtime_test.go

inspect:
	wasmer inspect build/dev_runtime.wasm