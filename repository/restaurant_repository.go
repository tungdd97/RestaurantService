package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"restaurant-service/database/mongo"
	"restaurant-service/entity"
)

type restaurantRepository struct {
	database   mongo.Database
	collection string
}

func NewRestaurantRepository(db mongo.Database, collection string) RestaurantRepository {
	return &restaurantRepository{
		database:   db,
		collection: collection,
	}
}

func (res *restaurantRepository) Create(c context.Context, restaurant *entity.Restaurant) error {
	collection := res.database.Collection(res.collection)

	_, err := collection.InsertOne(c, restaurant)

	return err
}

func (res *restaurantRepository) Fetch(c context.Context) ([]entity.Restaurant, error) {
	collection := res.database.Collection(res.collection)

	opts := options.Find().SetProjection(bson.M{})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var restaurants []entity.Restaurant

	err = cursor.All(c, &restaurants)
	if restaurants == nil {
		return []entity.Restaurant{}, err
	}

	return restaurants, err
}

func (res *restaurantRepository) FetchByCondition(c context.Context, condition bson.M) ([]entity.Restaurant, error) {
	collection := res.database.Collection(res.collection)

	opts := options.Find().SetProjection(bson.M{})
	cursor, err := collection.Find(c, condition, opts)

	if err != nil {
		return nil, err
	}

	var restaurants []entity.Restaurant

	err = cursor.All(c, &restaurants)
	if restaurants == nil {
		return []entity.Restaurant{}, err
	}

	return restaurants, err
}

func (res *restaurantRepository) GetByID(c context.Context, id string) (entity.Restaurant, error) {
	collection := res.database.Collection(res.collection)

	var restaurant entity.Restaurant

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return restaurant, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&restaurant)
	return restaurant, err
}

func (res *restaurantRepository) GetByCode(c context.Context, code string) (entity.Restaurant, error) {
	collection := res.database.Collection(res.collection)

	var restaurant entity.Restaurant
	err := collection.FindOne(c, bson.M{"code": code}).Decode(&restaurant)
	return restaurant, err
}

func (res *restaurantRepository) UpdateByID(c context.Context, id string, update interface{}) (int, error) {
	collection := res.database.Collection(res.collection)

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
