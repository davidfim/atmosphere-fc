package main

import (
	"log"
	"net/http"

	"github.com/davidfim/atmosphere-fc/api"
	"github.com/davidfim/atmosphere-fc/player"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=david password=qwertyuiop dbname=atmosphere port=5437"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&player.Player{})

	if err != nil {
		log.Fatal(err)
	}

	srv := api.NewServer(db)
	http.ListenAndServe(":8080", srv)
}
