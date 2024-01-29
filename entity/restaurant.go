package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionRestaurant = "restaurant"
)

type Restaurant struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Code        string             `bson:"code"`
	Description string             `bson:"description"`
	Rating      float32            `bson:"rating"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	CreatedBy   primitive.ObjectID `bson:"created_by"`
	UpdatedBy   primitive.ObjectID `bson:"updated_by"`
}
