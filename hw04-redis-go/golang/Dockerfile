# Dockerfile References: https://docs.docker.com/engine/reference/builder/
FROM golang
WORKDIR /go/src/app
COPY *.go ./
COPY go.mod .
COPY go.sum .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main .
CMD ["./main"]