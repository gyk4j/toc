FROM golang:alpine

WORKDIR /toc-backend

COPY go.mod .
COPY go.sum .
COPY cmd ./cmd/
COPY controllers ./controllers/
COPY models ./models/
COPY restapi ./restapi/
COPY services ./services/

RUN go mod tidy
RUN go build -o toc ./cmd/toc-server/main.go
RUN mkdir -p web
RUN tar -czvf toc-backend.tar.gz toc web/
