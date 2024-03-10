package routes

import (
	"github.com/gorilla/mux"

	"github.com/FinanceUN/Achievements/controllers"
)

func UserAchievementRoutes(r *mux.Router) {
	r.HandleFunc("/userAchievements", controllers.RegisterNewUser).Methods("POST")
	r.HandleFunc("/userAchievements", controllers.UpdateAchievementUserValue).Methods("PUT")
	r.HandleFunc("/userAchievements", controllers.GetUserAchievements).Methods("GET")
}
