FROM golang:latest as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go run scripts/generate/generate.go
RUN go build -o sthorer .

FROM golang:latest

COPY --from=builder /app/sthorer /
CMD ["/sthorer"]