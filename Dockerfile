FROM golang as builder

COPY main.go /go/build/main.go
COPY go.mod /go/build/go.mod
COPY go.sum /go/build/go.sum

WORKDIR /go/build/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .

FROM alpine:3.10

RUN mkdir /code
WORKDIR /code

COPY --from=builder /go/build/api /code/api

EXPOSE 8080

CMD ["/code/api"]
