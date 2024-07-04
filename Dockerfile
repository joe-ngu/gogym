FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o gogym .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/gogym /usr/local/bin/gogym

EXPOSE 8000

CMD ["gogym"]
