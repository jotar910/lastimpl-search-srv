FROM golang:1.17-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go get -u github.com/cosmtrek/air

EXPOSE 8081

ENTRYPOINT ["air"]