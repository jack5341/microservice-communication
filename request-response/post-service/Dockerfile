# builder image
FROM golang:1.17-alpine as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o /bin/deployment
FROM gcr.io/distroless/static:nonroot
WORKDIR /app/
COPY --from=builder /bin/deployment /app/deployment
ENTRYPOINT ["/app/deployment"]
