# Base image untuk Go
FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
COPY .env.production .env 

RUN go mod download

COPY . .

EXPOSE 3050

RUN go build -o binary

# Menjalankan binary
CMD ["/app/binary"]
