# builder container
FROM golang:1.14.1-alpine3.11 as builder

RUN apk update && apk add git

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go mod download
RUN go build -o main .

CMD ["/app/main"]


# distribution container
FROM alpine:3.11.5

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

# copy built assets from builder container
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]