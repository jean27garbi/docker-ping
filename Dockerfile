# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17-buster AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /docker-ping-roach /app/cmd/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /docker-ping-roach /docker-ping-roach

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-ping-roach"]