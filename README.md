# slack-timeline
manage your tasks

## Configurations

```sh
# Optional(port)
# port is defined by either environment variable or argument
export PORT=8000
```

## Getting started

```sh
# Run with specified port
go run main.go --port ${port}

# port is randomly assigned, if not specified
# except the case environment value - PORT is set
go run main.go

# health check returns json: {"messsage":"OK"}
curl <endpoint url>/health
```

## For developer

### Required

* go 1.12+

### Setup

1. mkdir %GOPATH%/src/github.com/mura123yasu
2. cd %GOPATH%/src/github.com/mura123yasu
3. git clone <this repository url>
4. cd <repository-name>
5. go mod download

### Unit Test

coming soon...

### Directory layout

coming soon...

### Build & Run

for Binary
```sh
# Build binary
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -a -installsuffix cgo -o ./bin/main ./main.go

# Run with specified port
./bin/main --port ${port}
```

for Docker
```sh
# Build
docker build -t slack-timeline .

# Run server
docker run --name slack-timeline-container -d -e "PORT=8000" -p 8000:8000 -t slack-timeline
```

## Deploy to Cloud Run
### from master branch
[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://deploy.cloud.run)

### from staging branch
[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://deploy.cloud.run/?revision=staging)
