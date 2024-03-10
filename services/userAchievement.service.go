package services

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/FinanceUN/Achievements/db"
	"github.com/FinanceUN/Achievements/models"
	"github.com/FinanceUN/Achievements/utils"
)

func RegisterNewUser(user utils.NewUser) (*mongo.InsertOneResult, error) {
	userExists, err := UserExists(user.UserID)
	if err != nil {
		return nil, err
	}

	if userExists {
		return nil, nil
	}

	achievementsTier0, err := GetAchievementsByTier(1)
	if err != nil {
		return nil, err
	}

	var userAchievements []models.UserAchievement

	for _, achievement := range achievementsTier0 {
		userAchievement := models.UserAchievement{
			AchievementID: achievement.ID,
			UserValue:     0,
			Reached:       false,
			CreatedAt:     primitive.NewDateTimeFromTime(time.Now()),
			UpdatedAt:     primitive.NewDateTimeFromTime(time.Now()),
		}

		userAchievements = append(userAchievements, userAchievement)
	}

	newUser := models.User{
		UserID:       user.UserID,
		Achievements: userAchievements,
		CreatedAt:    primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:    primitive.NewDateTimeFromTime(time.Now()),
	}

	insertResult, err := db.UserAchievementsCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func UserExists(userID string) (bool, error) {
	filter := bson.M{"user_id": userID}

	count, err := db.UserAchievementsCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func RequirementValueReached(achievementID string, userValue float64) (bool, error) {
	achievement, err := GetAchievement(achievementID)
	if err != nil {
		return false, err
	}

	if userValue >= achievement.RequirementValue {
		return true, nil
	}

	return false, nil
}

func UpdateAchievementUserValue(
	userUpdate utils.UserAchievementValueUpdate,
) (*mongo.UpdateResult, error) {
	achievementID, err := primitive.ObjectIDFromHex(userUpdate.AchievementID)
	if err != nil {
		return nil, err
	}

	reached, err := RequirementValueReached(userUpdate.AchievementID, userUpdate.Value)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"user_id": userUpdate.UserID, "achievements.achievement_id": achievementID}

	update := bson.M{
		"$set": bson.M{
			"achievements.$.user_value": userUpdate.Value,
			"achievements.$.reached":    reached,
			"achievements.$.updatedAt":  primitive.NewDateTimeFromTime(time.Now()),
		},
	}

	updateResult, err := db.UserAchievementsCollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func GetUserAchievements(userID string, options string) ([]models.UserAchievement, error) {
	filter := bson.M{
		"user_id": userID,
	}

	if options == "reached" {
		filter = bson.M{
			"user_id":              userID,
			"achievements.reached": true,
		}
	} else if options == "not_reached" {
		filter = bson.M{
			"user_id":              userID,
			"achievements.reached": false,
		}
	}

	var userAchievements []models.UserAchievement

	cursor, err := db.UserAchievementsCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user models.User
		cursor.Decode(&user)
		userAchievements = user.Achievements
	}

	return userAchievements, nil
}
