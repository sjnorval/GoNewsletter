package seed

import (
	"fmt"
	"strconv"
	"time"
	"github.com/sjnorval/fullstack/api/models"
	"github.com/sjnorval/fullstack/api/Utils"
)

func seedUsers() []User {
	for i := 0; i < 6; i++ {
		var name = "name" + strconv.Itoa(i)
		var surname = "surname" + strconv.Itoa(i)
		var users []User
		id, err := newUUID()

		if err != nil {
			fmt.Println("error: %v", err)
		}

		var user = User {
			ID:          id,
			Name:        name,
			Surname:     surname,
			Email:       name + surname + "@foo.co.za",
			LastUpdated: time.Now().String(),
		}

		users = append(users, user)

		return users
	}
}