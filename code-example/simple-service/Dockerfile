FROM golang:1.20.1

WORKDIR /app

# Build a minimal image instead of copying the whole project
COPY . .
RUN go mod download

RUN go build -o ./bin/simple-service ./cmd

EXPOSE 8888

CMD [ "./bin/simple-service" ]