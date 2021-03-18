FROM golang:1.15.10-alpine3.13

RUN apk add --no-cache git make gcc musl-dev
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make

EXPOSE 8080
ENTRYPOINT ["/app/bin/main"]
