package pubapisrv

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		dur := time.Now().Sub(start)
		log.Println(r.Method + " " + r.RequestURI + " " + dur.String())
	})
}

func forceHTTPS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proto := r.Header.Get("x-forwarded-proto")
		if proto == "http" || proto == "HTTP" {
			http.Redirect(w, r, fmt.Sprintf("https://%s%s", r.Host, r.URL), http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go"))
}

func New(addr string) (*http.Server, error) {
	r := mux.NewRouter()
	r.Use(logger)
	r.Use(forceHTTPS)
	r.HandleFunc("/", indexHandler)

	return &http.Server{
		Addr:         addr,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}, nil
}
