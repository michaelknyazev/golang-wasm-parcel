IMAGE_NAME=go_wasm_parcel

clean:
	rm -rf ./public/game.wasm ./dist ./.parcel-cache
compile:
	GOOS=js GOARCH=wasm go build -o ./public/game.wasm ./src/main.go
dev: clean compile
	parcel ./public/index.html ./public/assets/*
bundle:
	parcel build ./public/index.html ./public/assets/*

build: clean compile bundle

image:
	docker build -t ${IMAGE_NAME} .
serve: image
	docker run -p 3000:80 ${IMAGE_NAME}