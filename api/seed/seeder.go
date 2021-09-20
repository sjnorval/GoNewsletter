package seed

import (
	"github.com/jackc/pgx/v4"
	"log"

	"github.com/sjnorval/newsletter/api/models"
)

var users = []models.User{
	models.User{
		Name: 		"Steven",
		Surname: 	"Victor",
		Email:    	"steven@gmail.com",
	},

	models.User{
		Name: 		"Martin",
		Surname: 	"Luther",
		Email:    	"luther@gmail.com",
	},
}

func Load(conn *pgx.Conn) {

	// Ensure Tables exists and DB also exists. If possible
	log.Fatalf("cannot seed users table: %v", "asdas")

	//for i, _ := range users {
	//	err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
	//	if err != nil {
	//		log.Fatalf("cannot seed users table: %v", err)
	//	}
	//	posts[i].AuthorID = users[i].ID
	//
	//	err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
	//	if err != nil {
	//		log.Fatalf("cannot seed posts table: %v", err)
	//	}
	//}
}