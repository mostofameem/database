package handlers

import (
	"encoding/json"
	"fmt"
	"go_server_db/web/utils"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	var err error
	db, err = sqlx.Connect("postgres", "postgresql://root:admin@localhost:5432/root?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	// Decode JSON data from request body into a User struct
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		message := "Error converting user to JSON"
		utils.SendError(w, http.StatusNotAcceptable, message, newUser)
		return
	}

	// Insert the new user into the database
	err = InsertUser(newUser)
	if err != nil {
		log.Fatal("Error inserting user:", err)
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	}

	fmt.Println("User inserted successfully!")
	w.WriteHeader(http.StatusCreated)
}

// InsertUser inserts a new user into the database
func InsertUser(user User) error {
	_, err := db.Exec(`INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`, user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
