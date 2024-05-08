package handlers

import (
	"go_server_db/web/utils"
	"log"
	"net/http"
)

func ReadUsers() ([]User, error) {
	// Query the database to retrieve all users
	var users []User
	err := db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}
func Read(w http.ResponseWriter, r *http.Request) {
	users, err := ReadUsers()
	if err != nil {
		log.Fatal("Error fetching users:", err)
	}

	// Print the retrieved users
	//for _, user := range users {
	utils.SendJson(w, http.StatusAccepted, users)
	//}
}
