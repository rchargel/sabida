# syntax=docker/dockerfile:1

FROM golang:1.21.3

WORKDIR /app

# mv module info
COPY src/go.* ./

# install modules
RUN go mod download

# mv src code
COPY src/*.go ./
COPY src/dao ./dao
COPY src/handlers ./handlers
COPY src/migrations ./migrations
COPY src/models ./models
COPY src/static ./static
COPY src/templates ./templates

# compile source code
RUN CGO_ENABLED=0 GOOS=linux go build -o /sabida .

EXPOSE 8080

# Run
CMD ["/sabida"]
