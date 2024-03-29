# Build Stage
FROM golang:1.12.5 AS builder

ENV GO111MODULE on 
ENV GOOS linux 
ENV GOARCH amd64 
ENV CGO_ENABLED 0

WORKDIR /github.com/mura123yasu/slack-timeline
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ARG TARGETPATH="./main.go"
RUN go build -ldflags '-s -w' -a -installsuffix cgo -o /main ${TARGETPATH}

# Runtime Stage
FROM alpine:3.9.4
RUN apk add --no-cache ca-certificates
COPY --from=builder /main .

CMD ./main --host 0.0.0.0 --port ${PORT}

EXPOSE ${PORT}
