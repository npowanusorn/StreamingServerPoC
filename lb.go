package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/video/", func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		targetHost := replacePort(host, "9090")

		targetURL := "http://" + targetHost + r.URL.Path[len("/video"):]

		log.Printf("[LB] %s %s -> %s", r.Method, r.URL.Path, targetURL)

		http.Redirect(w, r, targetURL, http.StatusMovedPermanently)
	})

	log.Println("LB server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func replacePort(host, newPort string) string {
	if i := len(host); i > 0 {
		for j := len(host) - 1; j >= 0; j-- {
			if host[j] == ':' {
				return host[:j+1] + newPort
			}
		}
	}
	return host + ":" + newPort
}
