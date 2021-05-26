FROM golang:1.15

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]