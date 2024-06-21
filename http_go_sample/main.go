package main

import (
	"http_go_sample/webserver"
	"http_go_sample/webserver/dummy"
	"log"
	"net/http"
)

func main() {
	server := webserver.NewPlayerServer(dummy.NewInMemoryPlayerStore[webserver.Player]())
	log.Fatal(http.ListenAndServe(":5000", server))
}
