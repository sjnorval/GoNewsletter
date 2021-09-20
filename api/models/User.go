package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"html"
	"strings"
	"time"
)

type User struct {
	ID        	uint32    `json:"id"`
	Name  		string    `json:"name"`
	Surname  	string    `json:"surname"`
	Email     	string    `json:"email"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Surname = html.EscapeString(strings.TrimSpace(u.Surname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Surname == "" {
			return errors.New("Required Surname")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Surname == "" {
			return errors.New("Required Surname")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *pgx.Conn) (*User, error) {
	var err error
	var id uint32
	var sql = "INSERT INTO users (name, surname,email , createdat, updatedat)   VALUES ('$1', '$2', $3, DEFAULT, DEFAULT); SELECT currval(pg_get_serial_sequence('users','id'));"
	if err = db.QueryRow(context.Background(),sql, u.Name, u.Email, u.Email, u.CreatedAt, u.UpdatedAt).Scan(&id);  err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return &User{}, err
	}
	var user, _ = u.FindUserByID(db, id)

	return user, nil
}

func (u *User) FindAllUsers(db *pgx.Conn) (*[]User, error) {
	users := []User{}

	if rows, err := db.Query(context.Background(), "SELECT * FROM users"); err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return &[]User{}, err
	} else {
		defer rows.Close()
		var tmp User

		for rows.Next() {
			// Scan reads the values from the current row into tmp
			rows.Scan(&tmp)
			users = append(users, tmp)
			fmt.Printf("%+v\n", tmp)
		}
		if rows.Err() != nil {
			// if any error occurred while reading rows.
			fmt.Println("Error will reading user table: ", err)
			return &[]User{}, err
		}
	}

	return &users, nil
}

func (u *User) FindUserByID(db *pgx.Conn, uid uint32) (*User, error) {
	var err error
	var sql = "SELECT * FROM users WHERE id=$1"
	var foundUser *User
	if err = db.QueryRow(context.Background(),sql, uid).Scan(&foundUser);  err != nil {
		fmt.Println("User Not Found", err)
		return &User{}, err
	}

	return foundUser, err
}

func (u *User) UpdateAUser(db *pgx.Conn, uid uint32) (*User, error) {

	//was checking if the user exists
	//if foundUser, err = u.FindUserByID(db, uid);  err != nil {
	//	fmt.Println("User Not Found", err)
	//	return &User{}, err
	//}

	var foundUser *User
	var err error

	var sql = "UPDATE users SET name = '$1', surname = '$2', email = '$3', UpdatedAt = NOW() WHERE id=$2 ;"

	if _, err = db.Exec(context.Background(),sql, uid);  err != nil {
		fmt.Println("Update failed", err)
		return foundUser, err
	}

	if foundUser, err = u.FindUserByID(db, uid);  err != nil {
		fmt.Println("User Not Found", err)
		return &User{}, err
	}

	if err != nil {
		return &User{}, err
	}
	return foundUser, nil
}

//func (u *User) DeleteAUser(db *pgx.Conn, uid uint32) (int64, error) {
//
//	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
//
//	if db.Error != nil {
//		return 0, db.Error
//	}
//	return db.RowsAffected, nil
//}