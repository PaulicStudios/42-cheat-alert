FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN go build -o 42-cheat-alert

ENV TZ=Europe/Berlin

CMD ["./42-cheat-alert"]
