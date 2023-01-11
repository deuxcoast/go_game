package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/duexcoast/go_game"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("lets play poker")
	fmt.Println("Type: '{name} wins' to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening file %s, %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	game := poker.CLI{store, os.Stdin}
	game.PlayPoker()
}
