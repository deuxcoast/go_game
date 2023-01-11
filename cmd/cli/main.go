package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/duexcoast/go_game"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer close()

	game := poker.NewTexasHoldEm(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)

	fmt.Println("lets play poker")
	fmt.Println("Type: '{name} wins' to record a win")

	cli.PlayPoker()

}
