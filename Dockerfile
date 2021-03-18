FROM golang:1.15.10-alpine3.13

RUN apk add --no-cache make
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make

RUN go get -u github.com/cosmtrek/air

EXPOSE 8080
ENTRYPOINT ["/app/bin/main"]
