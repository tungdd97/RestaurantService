package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User model
type User struct {
	ID               primitive.ObjectID `bson:"_id"`
	Name             string             `bson:"name"`
	Password         string             `bson:"password"`
	TypeAccount      string             `bson:"type_account"`
	Email            *string            `bson:"email"`
	Address          string             `bson:"address"`
	PhoneNumber      string             `bson:"phone_number"`
	AccessTimeLatest time.Time          `bson:"access_time_latest"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
}

const (
	CollectionUser = "users"
)
