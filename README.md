# microservice-geoip
Microservice for detecting user country by client IP


### Installation 

```bash
go get 
```

### Run
```bash
go run main.go
```

### Test
```bash
curl "localhost:9090/country?ip=<YOUR_IP>"
```

### With Docker
```
docker build . -t geoip
```
To run geoip deamon 
```
docker run -d -p 9090:9090 --name geoip-instance --rm geoip
```