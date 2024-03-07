package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/FinanceUN/Achievements/db"
	"github.com/FinanceUN/Achievements/routes"
)

var router *mux.Router

func main() {
	db.DBConnection()

	router = mux.NewRouter()

    routes.IndexRoutes(router)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", router)
}
