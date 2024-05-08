package handlers

import (
	"fmt"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]

	_, er := db.Exec(`delete from users where id=$1;`, id)

	if er != nil {
		fmt.Fprintf(w, "Can not delete user")
		return
	}
	fmt.Fprintf(w, "Deleted Sucessfully")
}
