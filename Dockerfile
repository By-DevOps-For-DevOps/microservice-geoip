FROM golang:1.7.4
MAINTAINER kamoljan@gmail.com

ADD . /go/src/github.com/microservices-today/microservice-geoip

RUN go get github.com/microservices-today/microservice-geoip
RUN go install github.com/microservices-today/microservice-geoip

COPY GeoLite2-Country.mmdb /go/bin/microservice-geoip

ENTRYPOINT /go/bin/microservice-geoip

EXPOSE 9090