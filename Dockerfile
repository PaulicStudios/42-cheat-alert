FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN go build -o 42-cheat-alert

FROM gcr.io/distroless/base-debian11 AS build-release-stage
COPY --from=0 /app/42-cheat-alert /bin/42-cheat-alert
USER nonroot:nonroot
CMD ["/bin/42-cheat-alert"]
