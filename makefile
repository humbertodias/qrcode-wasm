serve:
	go run cmd/server/main.go

wasm:
	GOARCH=wasm GOOS=js go build -o assets/json.wasm cmd/wasm/main.go

wasm_exec:
	cp "`go env GOROOT`/misc/wasm/wasm_exec.js" assets/
