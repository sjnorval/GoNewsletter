package controllers

import (
	"fmt"
	"log"
	"net/http"
	"context"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/gin-gonic/gin"
	//"github.com/sjnorval/newsletter/api/models"
)

type Server struct {
	DB     *pgx.Conn
	Router *gin.Engine
}

func (server *Server) Initialize(DbUser, DbPassword, DbPort, DbHost, DbName, Sslmode string) {

	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", DbHost, DbPort, DbUser, DbName, DbPassword, Sslmode)

	conn, err := pgx.Connect(context.Background(), DBURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	if err != nil {
		fmt.Printf("Cannot connect to %s database", DbName)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", DbName)
	}

	server.Router = gin.Default()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}