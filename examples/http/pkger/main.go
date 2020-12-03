package main

import (
	"log"
	"net/http"

	"github.com/bingoohuang/pkger"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	dir := http.FileServer(pkger.Dir("/public"))
	return http.ListenAndServe(":3000", dir)
}
