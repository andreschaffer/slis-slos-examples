# build image
FROM golang:1.20.1-alpine as build

WORKDIR /build

COPY . .
RUN go mod download

RUN go build -o ./bin/simple-service ./cmd

# application image
FROM alpine:3.17.2
COPY --from=build /build/bin/simple-service .

EXPOSE 8888
EXPOSE 8889

CMD [ "./simple-service" ]