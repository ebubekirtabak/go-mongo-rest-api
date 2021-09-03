# Go Mongo Rest API

## Usage

Download dependencies

`go mod download`

Build API

`go build -o main .`

Run server

`./main`

## Running test

`go test ./...`


# Building Docker image

Build Docker image first

`docker build --tag member-api .`

Run Image

`docker run --publish 3000:3000 member-api`

## Documentation

[Documentation](https://documenter.getpostman.com/view/17301583/U16bx9bD)


