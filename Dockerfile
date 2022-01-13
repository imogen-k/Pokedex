FROM alpine:3.12

WORKDIR /app

FROM docker:dind
RUN apk add --no-cache go
RUN go version

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . ./

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["/pokedex"]