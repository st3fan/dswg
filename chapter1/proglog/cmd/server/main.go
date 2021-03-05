package main

import (
	"log"
	"github.com/st3fan/dswg/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
