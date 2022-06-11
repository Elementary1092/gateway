FROM golang:latest

ENV GO111MODULE=on

RUN git clone --depth 1 https://github.com/Elementary1092/gateway

WORKDIR /gateway

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o build/ cmd/main.go

EXPOSE 8080
CMD ["build/main"]