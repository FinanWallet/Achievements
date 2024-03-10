package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/FinanceUN/Achievements/db"
	"github.com/FinanceUN/Achievements/routes"
)

func main() {
	client := db.DBConnection()

	router := mux.NewRouter()

	routes.IndexRoutes(router)
	routes.AchievementsRoutes(router)
    routes.UserAchievementRoutes(router)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", router)

    if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
