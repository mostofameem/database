package main

import (
	"go_server_db/web"
	"log"
	"net/http"
)

func main() {

	mux := web.StartServer()
	log.Printf("Server Running on port 30000")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
