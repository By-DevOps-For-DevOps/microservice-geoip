# microservice-geoip

Microservice for Detecting Country by IP Address

## Installation

```shell
go get 
```

## Try locally

```shell
export PORT=9090
go run main.go
```

### Get Country by IP
```shell
curl localhost:9090/country?ip=<YOUR_IP>
```

### Get Outbound IP
```shell
curl localhost:9090/outbound
```

## Run on Docker
```shell
docker build . -t geoip
```
To run geoip deamon 
```shell
docker run -d -p 9090:9090 --name geoip-instance --rm geoip
```

## Testing
Run the following command
```shell
go test -v ./...
```
