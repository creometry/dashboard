FROM golang:1.18 AS build

WORKDIR /go/src/github.com/Creometry/resources-service

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


FROM alpine:latest

WORKDIR /app

# these copies are just to make the image work, they are not needed when we deploy with kubernetes
COPY ./kubeconfig.yaml /app/kubeconfig.yaml
COPY ./.env /app/.env

COPY --from=build /go/src/github.com/Creometry/resources-service/app /app

EXPOSE 3002

CMD ["./app"]

