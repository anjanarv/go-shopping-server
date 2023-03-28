FROM golang:1.20.2-buster as builder

WORKDIR /app
#
COPY go.mod ./
COPY go.sum ./

#COPY go.mod ./
#RUN go mod download
RUN go mod download

COPY api/*.go ./

# Copy local code to the container image.
#COPY . ./
#COPY . .

RUN go build -o /sampleserver
#
#FROM scratch as final
#COPY --from=builder /app/sampleserver /sampleserver

EXPOSE 8080

ENTRYPOINT [ "/sampleserver" ]