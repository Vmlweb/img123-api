FROM golang:1.14 AS builder

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o img123 main.go

FROM gcr.io/distroless/base-debian10

COPY --from=builder /img123 .

ENTRYPOINT ["/img123"]