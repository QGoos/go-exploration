package main

import (
	"http_go_sample/dummy"
	"http_go_sample/webserver"
	"log"
	"net/http"
)

func main() {
	server := &webserver.PlayerServer{dummy.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
