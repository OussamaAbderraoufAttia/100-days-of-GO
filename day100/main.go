package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"sync"
)

var (
	urls  = make(map[string]string)
	mutex sync.Mutex
)

func shorten(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("url")
	if longURL == "" {
		http.Error(w, "Url missing", 400)
		return
	}

	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(longURL)))[:6]

	mutex.Lock()
	urls[hash] = longURL
	mutex.Unlock()

	fmt.Fprintf(w, "Shortened! Use: http://localhost:8080/%s", hash)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	mutex.Lock()
	longURL, ok := urls[code]
	mutex.Unlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/shorten", shorten)
	http.HandleFunc("/", redirect)

	fmt.Println("URL Shortener running on :8080")
	http.ListenAndServe(":8080", nil)
}
