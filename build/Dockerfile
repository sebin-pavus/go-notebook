# syntax=docker/dockerfile:1

FROM golang:1.19-alpine
# We'll put the app in /app in the container
WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# Copy everything in
COPY . .

RUN go build -o /main ./cmd/server/main.go

EXPOSE 8080

CMD [ "/main" ]