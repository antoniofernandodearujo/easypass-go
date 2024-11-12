FROM golang:1.22

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o easypass-go ./cmd/server

CMD ["./easypass-go"]

EXPOSE 8080