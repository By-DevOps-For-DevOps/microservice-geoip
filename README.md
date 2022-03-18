# microservice-geoip

Microservice for detecting user country by client IP

### Installation 

```shell
go get 
```

### Run

```shell
go run main.go
```

### Test

```shell
curl "localhost:9090/country?ip=<YOUR_IP>"
```

#### With Docker
```shell
docker build . -t geoip
```
To run geoip deamon 
```shell
docker run -d -p 9090:9090 --name geoip-instance --rm geoip
```
