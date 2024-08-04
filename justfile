run: css client
	go run .

css:
	tailwindcss -i input.css -o static/style.css

client:
	GOOS=js GOARCH=wasm go build -o static/client.wasm ./client
	install "$(go env GOROOT)/misc/wasm/wasm_exec.js" static/
