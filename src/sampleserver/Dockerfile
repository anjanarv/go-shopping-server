FROM golang:1.20-buster

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /sampleserver

EXPOSE 8080

CMD [ "/sampleserver" ]