FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /fileserver


FROM scratch
WORKDIR /
COPY --from=0 /fileserver /

EXPOSE 80

ENTRYPOINT ["/fileserver"]

