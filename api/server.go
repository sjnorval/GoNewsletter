package api

import (
	//"fmt"
	"log"
	"os"

	"github.com/sjnorval/newsletter/api/controllers"
	"github.com/sjnorval/newsletter/api/seed"
)

var server = controllers.Server{}

func init() {
	// loads values from .env into the system
	if err := os.Environ(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":8080")
}