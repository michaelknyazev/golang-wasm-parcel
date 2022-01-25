clean:
	rm -rf ./public/game.wasm ./dist ./.parcel-cache
compile:
	GOOS=js GOARCH=wasm go build -o ./public/game.wasm ./src/main.go
dev: clean compile
	parcel ./public/index.html ./public/assets/*