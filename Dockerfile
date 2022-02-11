FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go get github.com/githubnemo/CompileDaemon


ENTRYPOINT CompileDaemon --build="go build server.go" --command=./server


# FROM golang:alpine

# # Required because go requires gcc to build
# RUN apk add build-base

# RUN apk add inotify-tools

# WORKDIR $GOPATH/src/golang-api-gin

# ADD . ./

# COPY go.mod ./

# COPY go.sum ./

# RUN go mod download

# COPY . .

# ENV PORT 8080

# RUN go build -o golang-api-gin
# EXPOSE 8080

# ENTRYPOINT  ["./golang-api-gin"]



