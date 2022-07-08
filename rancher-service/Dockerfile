FROM golang:1.18 AS build

WORKDIR /go/src/github.com/Creometry/rancher-service

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


FROM alpine:latest

WORKDIR /app

COPY ./.env /app/.env

COPY --from=build /go/src/github.com/Creometry/rancher-service/app /app

EXPOSE 3001

CMD ["./app"]

