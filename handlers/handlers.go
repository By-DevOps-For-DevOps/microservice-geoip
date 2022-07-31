package handlers

import (
	"github.com/microservices-today/microservice-geoip/db"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

func Country(w http.ResponseWriter, r *http.Request) {
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
	record, err := db.Reader.Country(ip)
	if err != nil {
		log.Fatal(err)
	}
	if ip != nil {
		io.WriteString(w, record.Country.IsoCode)
	} else {
		http.Error(w, "IP not found", http.StatusNotFound)
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `Don't worry, I'm healthy!`)
}

func Ip(w http.ResponseWriter, r *http.Request) {
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

func Outbound(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://ifconfig.me")
	if err != nil {
		http.Error(w, "Something is wrong with ifconfig.me", http.StatusBadRequest)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	w.Write(body)
}
