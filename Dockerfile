FROM golang:1.17 as builder
WORKDIR /app
COPY . .
RUN make compile

FROM node:alpine as bundler
RUN apk add make
RUN npm i -g parcel
WORKDIR /app/
COPY . .
COPY --from=builder /app/public/game.wasm ./public
RUN make bundle

FROM nginx:alpine
WORKDIR /usr/share/nginx/html
COPY --from=bundler /app/dist .

EXPOSE 80