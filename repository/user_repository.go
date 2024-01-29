package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"restaurant-service/database/mongo"
	"restaurant-service/entity"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *entity.User) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]entity.User, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []entity.User

	err = cursor.All(c, &users)
	if users == nil {
		return []entity.User{}, err
	}

	return users, err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (entity.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user entity.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (entity.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user entity.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}

func (ur *userRepository) UpdateByID(c context.Context, id string, update interface{}) (int, error) {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	resultUpdate, err := collection.UpdateOne(c, bson.M{"_id": idHex}, update)
	if err != nil {
		return 0, err
	}
	return int(resultUpdate.MatchedCount), err
}
