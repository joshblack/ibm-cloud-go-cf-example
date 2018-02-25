package main

import (
	"log"
	"os"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
	"github.com/joshblack/ibm-cloud-go-cf-example/internal/pubapi/pubapisrv"
)

type config struct {
	addr string
	port string
	host string
}

func main() {
	var cfg config

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	cfg.port = port

	if cfenv.IsRunningOnCF() {
		cfg.host = "0.0.0.0"
	} else {
		cfg.host = "localhost"
	}
	cfg.addr = cfg.host + ":" + cfg.port

	srv, err := pubapisrv.New(cfg.addr)
	if err != nil {
		log.Fatalf("error setting up the server: %v", err)
	}

	log.Fatal(srv.ListenAndServe())
}
