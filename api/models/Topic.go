package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"html"
	"strings"
)

type Topic struct {
	ID        	uint32 `json:"id"`
	Description string `json:"description"`
}

func (t *Topic) Prepare() {
	t.ID = 0
	t.Description = html.EscapeString(strings.TrimSpace(t.Description))
}

func (t *Topic) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if t.Description == "" {
			return errors.New("Required Description")
		}
		return nil

	default:
		if t.Description == "" {
			return errors.New("Required Description")
		}
		return nil
	}
}

func (t *Topic) SaveTopic(db *pgx.Conn) (*Topic, error) {
	var err error
	var id uint32
	var sql = "INSERT INTO topics (description)   VALUES ('$1'); SELECT currval(pg_get_serial_sequence('topics','id'));"
	if err = db.QueryRow(context.Background(),sql, t.Description).Scan(&id);  err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return &Topic{}, err
	}
	var topic, _ = t.FindTopicByID(db, id)

	return topic, nil
}

func (t *Topic) FindTopicByID(db *pgx.Conn, uid uint32) (*Topic, error) {
	var err error
	var sql = "SELECT * FROM users WHERE id=$1"
	var foundTopic *Topic
	if err = db.QueryRow(context.Background(),sql, uid).Scan(&foundTopic);  err != nil {
		fmt.Println("User Not Found", err)
		return &Topic{}, err
	}

	return foundTopic, err
}

func (t *Topic) FindAllTopics(db *pgx.Conn) (*[]Topic, error) {
	topics := []Topic{}

	if rows, err := db.Query(context.Background(), "SELECT * FROM topics"); err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return &[]Topic{}, err
	} else {
		defer rows.Close()
		var tmp Topic

		for rows.Next() {
			// Scan reads the values from the current row into tmp
			rows.Scan(&tmp)
			topics = append(topics, tmp)
			fmt.Printf("%+v\n", tmp)
		}
		if rows.Err() != nil {
			// if any error occurred while reading rows.
			fmt.Println("Error will reading topics table: ", err)
			return &[]Topic{}, err
		}
	}

	return &topics, nil
}