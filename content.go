package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./content"))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[CONTENT] %s %s", r.Method, r.URL.Path)

		// CORS (useful for debugging tools / browsers)
		w.Header().Set("Access-Control-Allow-Origin", "*")

		fs.ServeHTTP(w, r)
	})

	log.Println("Content server running on http://localhost:9090")
	err := http.ListenAndServe(":9090", handler)
	if err != nil {
		log.Fatal(err)
	}
}
