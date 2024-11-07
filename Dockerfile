FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/app/main.go

FROM alpine:latest

RUN apk --update add ca-certificates curl && rm -rf /var/cache/apk/* && apk add --no-cache curl

WORKDIR /app

EXPOSE 8080

COPY --from=build /app/main /app/.env ./

COPY --from=build /app/pkg/template /app

CMD ["./main"]