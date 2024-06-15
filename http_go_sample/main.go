package main

import (
	"http_go_sample/dummy"
	"http_go_sample/webserver"
	"log"
	"net/http"
)

func main() {
	error_fix := &dummy.InMemoryPlayerStore{}
	server := &webserver.PlayerServer{error_fix}
	log.Fatal(http.ListenAndServe(":5000", server))
}
