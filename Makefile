build_wasm:
# TODO use https://github.com/radkomih/tinygo
# tinygo build \
# 	-wasm-abi=generic \
# 	-scheduler=none \
# 	-gc=none \
# 	-o=build/dev_runtime.wasm \
# 	runtime/runtime.go
	wat2wasm -o build/dev_runtime.wasm build/dev_runtime.wat 

inspect:
	wasmer inspect build/dev_runtime.wasm

validate:
	wasm-validate -v build/dev_runtime.wasm
	
run:
	go run main.go

test:
	go test -v runtime/runtime_test.go