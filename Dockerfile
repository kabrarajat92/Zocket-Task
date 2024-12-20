FROM golang:1.19-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
CMD ["./main"]
EXPOSE 8080
