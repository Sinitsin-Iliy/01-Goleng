FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download && go build -o app ./cmd/main.go

FROM alpine
COPY --from=builder /app/app /app/
CMD [ "/app/app" ]
EXPOSE 8000