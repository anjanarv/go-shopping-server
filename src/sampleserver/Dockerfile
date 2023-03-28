FROM golang:1.20.2-buster as builder

WORKDIR /api-files
#
#COPY go.mod ./
#COPY go.sum ./

COPY go.mod ./
RUN go mod download

# Copy local code to the container image.
COPY . ./
#COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o ./sampleserver

#RUN go mod download

#COPY *.go ./

#RUN go build -o /sampleserver

FROM scratch as final
COPY --from=builder /api-files/sampleserver /sampleserver

EXPOSE 8080

ENTRYPOINT [ "/sampleserver" ]