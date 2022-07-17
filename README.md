Bug reproduction

```
tinygo build -target wasi -wasm-abi generic -o module1.wasm ./impl/wasi/runtime/runtime.go
tinygo build -target wasi -wasm-abi generic -o module2.wasm main.go
go run ./cmd/wasibug/main.go
```
