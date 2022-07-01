# build stage
FROM golang:1.18 AS build-env
ADD . /src
WORKDIR /src
RUN go get
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main

# final stage
FROM alpine
COPY --from=build-env /src/main /
ENTRYPOINT /main
