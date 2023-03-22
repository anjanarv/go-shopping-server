FROM golang:1.20

WORKDIR /app/go-sample-app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o ./out/go-sample-app .

EXPOSE 8080

CMD [ "./out/go-sample-app" ]