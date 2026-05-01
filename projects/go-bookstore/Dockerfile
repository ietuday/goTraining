FROM golang:1.26-alpine AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bookstore-api .

FROM alpine:3.22

WORKDIR /app

COPY --from=build /bookstore-api /bookstore-api

EXPOSE 8000

ENTRYPOINT ["/bookstore-api"]
