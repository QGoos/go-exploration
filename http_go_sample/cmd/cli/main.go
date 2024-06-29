package main

import (
	"fmt"
	"http_go_sample/poker"
	"http_go_sample/webserver"
	"log"
	"os"
)

const dbFileName = "../game.db.json"

// var dummySpyAlerter = &poker.SpyBlindAlerter{}

func main() {
	store, closer, err := webserver.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	fmt.Println("Let's play poker!")
	fmt.Println("Type {Name} wins to record a win.")

	game := poker.NewGame(poker.BlindAlerterFunc(poker.Alerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
