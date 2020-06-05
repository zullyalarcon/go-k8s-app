FROM golang:alpine as builder
LABEL maintainer="Zully Alarcon <zullyalarcon@gmail.com>"
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main
CMD ["/app/main"]