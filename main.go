package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/oschwald/geoip2-golang"
	"strings"
)

var db *geoip2.Reader

func main() {
	var err error
	db, err = geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	http.HandleFunc("/country", country) // e.g. /country?ip=81.2.69.142
	http.HandleFunc("/ip", ip)
	http.HandleFunc("/health", health)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}

func health(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `Don't worry, I'm healthy!`)
}

func ip(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "GET only", http.StatusMethodNotAllowed)
		return
	}
	// Assuming format is as expected
	ips := strings.Split(r.Header.Get("X-Forwarded-For"), ", ")
	for _, ip := range ips {
		w.Write([]byte(ip + "\n"))
	}
}

func country(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "GET only", http.StatusMethodNotAllowed)
		return
	}
	qIP := r.URL.Query().Get("ip")
	if len(qIP) == 0 {
		http.Error(w, "IP is missing", http.StatusBadRequest)
		return
	}
	ip := net.ParseIP(qIP)
	record, err := db.Country(ip)
	if err != nil {
		log.Fatal(err)
	}
	if ip != nil {
		io.WriteString(w, record.Country.IsoCode)
	} else {
		http.Error(w, "IP not found", http.StatusNotFound)
	}
}
