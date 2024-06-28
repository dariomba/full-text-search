FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/full-text-search ./src/cmd

FROM gcr.io/distroless/static-debian11

WORKDIR /root/

COPY --from=builder /app/full-text-search .

COPY --from=builder /app/datasets ./datasets

EXPOSE 8080

CMD ["./full-text-search"]
