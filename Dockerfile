# Base build image
FROM golang:1.12
WORKDIR /go/src/bitbucket.org/wunzi/woinder-backend
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./main.go
#CMD ["./main"]
RUN go build ./main.go
#CMD ["./main"]
#RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/weaviate-server
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/src/bitbucket.org/wunzi/woinder-backend
#CMD ["./main"]
EXPOSE 8080
