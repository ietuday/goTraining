FROM golang:1.26.2-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY . .
RUN go mod tidy
RUN go build -o /api .

FROM alpine:3.22

WORKDIR /app
COPY --from=build /api /app/api

EXPOSE 8000
CMD ["/app/api"]
