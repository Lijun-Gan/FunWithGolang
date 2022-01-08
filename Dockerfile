FROM golang:1.17.3 

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]

# docker build -t goo_book .
# docker run -p 5000:5000 -it go_book 

#build stage
# FROM golang:alpine as builder
# RUN apk update && apk add --no-cache git
# RUN mkdir /build 
# ADD . /build/
# WORKDIR /build
# RUN go get -d -v
# RUN go build -o client .

# # Stage 2
# FROM alpine
# RUN adduser -S -D -H -h /app appuser
# USER appuser
# COPY --from=builder /build/ /app/
# WORKDIR /app
# CMD ["./client"]