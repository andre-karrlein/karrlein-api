FROM golang as builder

COPY main.go /go/build/main.go
WORKDIR /go/build/

RUN go install

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .

FROM alpine:3.10

RUN mkdir /code
WORKDIR /code

COPY --from=builder /go/build/api /code/api

EXPOSE 8080

CMD ["/code/api"]
