FROM golang:1.20 as builder

WORKDIR /app
#
#COPY go.mod ./
#COPY go.sum ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o ./sampleserver ./app/./...

#RUN go mod download

#COPY *.go ./

#RUN go build -o /sampleserver

FROM scratch as final
COPY --from=builder /app/sampleserver /sampleserver

EXPOSE 8080

ENTRYPOINT [ "/sampleserver" ]