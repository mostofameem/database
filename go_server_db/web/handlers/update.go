package handlers

import (
	"encoding/json"
	"fmt"
	"go_server_db/web/utils"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		message := "Error converting user to JSON"
		utils.SendError(w, http.StatusNotAcceptable, message, user)
		return
	}
	_, er := db.Exec(`UPDATE users SET name = $1 WHERE id = $2`, user.Name, id)

	if er != nil {
		message := "Can not update the value"
		utils.SendError(w, http.StatusNotAcceptable, message, user)
	}
	fmt.Fprintf(w, "Update Sucessfully")
}
