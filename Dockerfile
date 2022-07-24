FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY GeoLite2-Country.mmdb ./

RUN go build -o ./geoip

EXPOSE 9090
CMD ["/app/geoip"]
