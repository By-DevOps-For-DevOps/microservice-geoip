# microservice-geoip

Microservice for detecting user country by client IP

### Installation 

```shell
go get 
```

### Run

```shell
export PORT=9090
go run main.go
```

### Test

Country by IP
```shell
curl localhost:9090/country?ip=<YOUR_IP>
```

Outbound IP
```shell
curl localhost:9090/outbound
```

#### With Docker
```shell
docker build . -t geoip
```
To run geoip deamon 
```shell
docker run -d -p 9090:9090 --name geoip-instance --rm geoip
```
