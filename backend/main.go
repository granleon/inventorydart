package main

import (
	"net/http"

	"fmt"
	"os"

	"github.com/plleo/inventory/backend/app"
	"github.com/plleo/inventory/backend/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(app.JwtAuthentication)

	r.HandleFunc("/api/v1/user/new", controllers.CreateAccount).Methods("POST")
	r.HandleFunc("/api/v1/user/login", controllers.AuthenticateAccount).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}

	fmt.Println(port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println(err)
	}
}
