.PHONY: package generate compile

package:
	wkg wit build

generate:
	go tool wit-bindgen-go generate --world examples --out internal ./enclave:examples-go.wasm

compile:
	tinygo build -target=wasip2 -o examples-go.wasm --wit-package enclave:examples-go.wasm --wit-world examples-go main.go 2>&1
