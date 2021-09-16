package api

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	s "strings"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	"github.com/sjnorval/fullstack/api/models"
	"github.com/sjnorval/fullstack/api/Utils"
)

var users []User
var topics []Topic

func mayRegisterNewUser(user *User, users []User) bool {

	if len(users) < 1 {
		return true
	}

	sort.Slice(users, func(i, j int) bool {
		return s.ToLower(users[i].Email) <= s.ToLower(users[j].Email)
	})

	idx := sort.Search(len(users), func(i int) bool {
		return s.ToLower(string(users[i].Email)) >= s.ToLower(user.Email)
	})

	return s.ToLower(users[idx].Email) == s.ToLower(user.Email)
}

func maySubscribeToTopic(user *User, topic *Topic) bool {

	if len(user.Topics) < 1 {
		return true
	}

	sort.Slice(user.Topics, func(i, j int) bool {
		return user.Topics[i].ID <= topics[j].ID
	})

	idx := sort.Search(len(user.Topics), func(i int) bool {
		return string(user.Topics[i].ID) >= topic.ID
	})

	return user.Topics[idx].ID == topic.ID
}

func userById(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func registerUser(c *gin.Context) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		return
	}

	//check if user is unique

	//do a to lower on email

	users = append(users, user)
	c.IndentedJSON(http.StatusCreated, user)
}

func updateUser(c *gin.Context) {

}

func getTopics(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, topics)
}

func disconnectUserFromTopic(c *gin.Context) {

}

func connectUserToTopic(c *gin.Context) {
	id := c.Param("id")

	var topic Topic
	if err := c.BindJSON(&topic); err != nil {
		return //Not valid
	}

	for _, user := range users {
		if user.ID == id {

			if !maySubscribeToTopic(&user, &topic) {
				c.IndentedJSON(http.StatusConflict, gin.H{"message": "user already has topic listed"})
				c.IndentedJSON(http.StatusConflict, gin.H{"action": "/users/" + user.ID})
			}

			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func getUserTopics(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user.Topics)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func seedTopics() {
	var techId = "a6a4b8d8-fd91-45f3-b046-a863565c45dc"
	var musicId = "04ee280c-1b71-4fd9-bbb6-6f77842be301"
	var healthId = "556cee5d-0242-4ed2-8e75-760b503d83f0"

	topics = append(topics,
		Topic{ID: techId, Description: "Technology"},
		Topic{ID: musicId, Description: "Music"},
		Topic{ID: healthId, Description: "Health"},
	)
}

//Copied from https://play.golang.org/p/4FkNSiUDMg
//Not gonna design this from scratch. Only had a week to learn this language and sorry if this dissapoints.
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}



func Run() {
	seedUsers()
	seedTopics()

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/users/:id", userById)
	router.GET("/users/user/topics/:id", getUserTopics)
	router.POST("/users/:id", updateUser)
	router.POST("/users/register", registerUser)
	router.GET("/topics", getTopics)
	router.POST("/topics/user/:id", connectUserToTopic)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
