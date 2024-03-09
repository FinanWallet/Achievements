package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/FinanceUN/Achievements/models"
	"github.com/FinanceUN/Achievements/db"
)

func CreateAchievement(achievement models.Achievement) (*mongo.InsertOneResult, error) {
	insertResult, err := db.AchievementsCollection.InsertOne(context.TODO(), achievement)
	return insertResult, err
}
