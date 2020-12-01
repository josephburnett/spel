.goals = build
.PHONY : build

build :
	cp $$GOROOT/misc/wasm/wasm_exec.js resources/
	GOOS=js GOARCH=wasm go build -o resources/spel.wasm ./cmd/ui/main.go
