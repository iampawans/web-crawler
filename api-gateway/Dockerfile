FROM golang:1.18-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o api-gateway
EXPOSE 8082
CMD ["./api-gateway"]