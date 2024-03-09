package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAchievement struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"            json:"_id"`
	UserID        primitive.ObjectID `bson:"user_id,omitempty"        json:"user_id"`
	AchievementID primitive.ObjectID `bson:"achievement_id,omitempty" json:"achievement_id"`
	UserValue     int                `bson:"user_value,omitempty"     json:"user_value"`
	Reached       bool               `bson:"reached,omitempty"        json:"reached"`
	CreatedAt     primitive.DateTime `bson:"createdAt,omitempty"      json:"createdAt"`
	UpdatedAt     primitive.DateTime `bson:"updatedAt,omitempty"      json:"updatedAt"`
}
