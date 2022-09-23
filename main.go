package main

import (
	"fmt"
	"github.com/microservices-today/microservice-geoip/db"
	"github.com/microservices-today/microservice-geoip/handlers"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net/http"
	"os"
)

func main() {
	var err error
	db.Reader, err = geoip2.Open("./db/GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Reader.Close()

	http.HandleFunc("/country", handlers.Country) // e.g. /country?ip=81.2.69.142
	http.HandleFunc("/ip", handlers.Ip)
	http.HandleFunc("/outbound", handlers.Outbound)
	http.HandleFunc("/health", handlers.Health)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
