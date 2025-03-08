package main

import (
	"fmt"
	"zigzag-core/route"

	"github.com/LabStars/selpo-common/crypto"
	"github.com/LabStars/selpo-common/db"
)

func main() {
	dbConfig, err := db.ReadJSONFromFile("./config/dbSettings.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	key, err := db.ReadUserInput()
	if err != nil {
		fmt.Println("Failed to read from console:", err)
		return
	}

	dbPassword, err := crypto.Decrypt(dbConfig.Password, key)
	if err != nil {
		fmt.Printf("Decryption failed: %v", err)
		return
	}

	db, err := db.ConnectDatabase(&dbConfig, dbPassword)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	route.StartServer(db, 8080)
}
