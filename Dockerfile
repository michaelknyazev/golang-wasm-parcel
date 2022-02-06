FROM golang:1.17 as builder
WORKDIR /app/
COPY go.mod go.sum .
COPY src/ .
RUN GOOS=js GOARCH=wasm go build -o game.wasm main.go

FROM node:alpine as bundler
RUN npm i -g parcel
WORKDIR /public/
COPY --from=builder /app/game.wasm .
COPY public/ .
RUN parcel build index.html assets/*

FROM nginx:alpine
WORKDIR /usr/share/nginx/html
COPY --from=bundler /public/dist .

EXPOSE 80