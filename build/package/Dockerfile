FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main ./cmd/stocks-forwarder
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./main"]
