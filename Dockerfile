# First Stage
FROM golang:alpine as build

RUN apk add --no-cache git

WORKDIR /go/src/github.com/yurichandra/shrt

ADD . .

RUN go get -u -v github.com/golang/dep/cmd/dep
RUN dep ensure -v

RUN CGO_ENABLED=0 go build

# Second Stage
FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/github.com/yurichandra/shrt .

CMD [ "./shrt", "serve"]
