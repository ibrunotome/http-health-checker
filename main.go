package main

import (
	"net/http"
	hc "pkg/healthchecker"
)

func main() {
	http.HandleFunc("/", hc.HealthCheck)
	http.ListenAndServe(":80", nil)
}
