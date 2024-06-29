package main

import (
	"http_go_sample/poker"
	"http_go_sample/webserver"
	"log"
	"net/http"
)

const dbFileName = "../game.db.json"

func main() {
	store, closer, err := webserver.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	game := poker.NewGame(poker.BlindAlerterFunc(poker.Alerter), store)

	server, err := webserver.NewPlayerServer(store, game)

	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Not good man: could not listen on port 5000 %v", err)
	}

	// db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	// if err != nil {
	// 	log.Fatalf("Oh shit! problem openint %s %v", dbFileName, err)
	// }

	// store, err := webserver.NewFileSystemPlayerStore(db)

	// if err != nil {
	// 	log.Fatalf("problem creating file system player store, %v ", err)
	// }

	// server := webserver.NewPlayerServer(store)

	// server := webserver.NewPlayerServer(dummy.NewInMemoryPlayerStore[webserver.League]())
	// if err := http.ListenAndServe(":5000", server); err != nil {
	// 	log.Fatalf("Not good man: could not listen on port 5000 %v", err)
	// }
}
