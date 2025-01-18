FROM golang:1.18-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o data-service
EXPOSE 8081
CMD ["./data-service"]