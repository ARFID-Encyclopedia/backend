package main

import (
	"encoding/json"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/sdomino/scribble"
)

func initDB() *scribble.Driver {
	db, err := scribble.New("./data", nil)
	if err != nil {
		log.Fatal("Error initialising the database", err)
	}
	return db
}

func getAllUsers() []user {
	users := []user{}
	records, err := db.ReadAll("users")
	if err != nil {
		log.Println("Error", err)
	}
	for _, u := range records {
		user := user{}
		if err := json.Unmarshal(u, &user); err != nil {
			log.Println("Unable to unmarshal user", err)
		}
		users = append(users, user)
	}
	return users
}

func addTestUser() {
	user := user{
		ID:           uuid.Must(uuid.NewV4()).String(),
		Username:     "test",
		PasswordHash: "",
		AccessLevel:  1,
	}
	if err := db.Write("users", user.ID, user); err != nil {
		fmt.Println("Error", err)
	}
}

func genUserMap() {
	userMap = make(map[string]string)
	users := getAllUsers()
	for _, u := range users {
		userMap[u.Username] = u.ID
	}
}

func writeUserToDB(user user) error {
	if err := db.Write("users", user.ID, user); err != nil {
		return err
	}
	return nil
}

func writeFoodToDB(food food) error {
	if err := db.Write("foods", food.Name, food); err != nil {
		return err
	}
	return nil
}
