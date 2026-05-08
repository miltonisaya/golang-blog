package main

import (
	"blog/inits"
	"blog/models"
	"log"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	err := inits.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal(err)
		return
	}
}
