.goals = build serve
.PHONY : build serve

build :
	cp $$GOROOT/misc/wasm/wasm_exec.js resources/
	GOOS=js GOARCH=wasm go build -o resources/spel.wasm ./cmd/ui/main.go

serve : build
	cd resources ; python -m SimpleHTTPServer 8000
