package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID     string `json:"id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age    int    `json:"age" bson:"age"`
}

// StoreUserData saves the given u user data map into a file. Always overrides everything
func StoreUserData(u map[string]User) {
	file, err := os.Create("db.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	json.NewEncoder(file).Encode(u) //from variable
}

// LoadUserData returns the map corresponding to the user data from the file
func LoadUserData() map[string]User {
	u := make(map[string]User)

	file, err := os.Open("db.json")
	if err != nil {
		fmt.Println(err)
		return u
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&u) //into variable
	if err != nil {
		fmt.Println(err)
	}
	return u
}
