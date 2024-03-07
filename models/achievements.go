package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Achievement struct {
    ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Title       string             `json:"title,omitempty" bson:"title,omitempty"`
    Description string             `json:"description,omitempty" bson:"description,omitempty"`
    RequirementValue int           `json:"requirementValue,omitempty" bson:"requirementValue,omitempty"`
    CreatedAt   primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
    UpdatedAt   primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
