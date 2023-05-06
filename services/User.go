package services

import (
	"app-user/models"
	"app-user/utils"
	"context"
	"crypto/md5"
	"encoding/hex"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	Data []*models.UserModel
}

func InitUserService() *UserService {
	return &UserService{}
}

func (userService UserService) CreateUserService(ctx context.Context, mongoCollection *mongo.Collection, schema models.UserModel) ([]*models.UserModel, error) {
	hashId, err := bson.Marshal(schema)
	if err != nil {
		return nil, err
	}
	hash := md5.Sum(hashId)
	schema.Id = hex.EncodeToString(hash[:])
	schema.Password, err = utils.HashPassword([]byte(schema.Password))
	if err != nil {
		return nil, err
	}
	_, err = mongoCollection.InsertOne(ctx, schema)
	if err != nil {
		return nil, err
	}
	userService.Data = append(userService.Data, &schema)
	return userService.Data, nil
}

func (userService UserService) GetUserService(ctx context.Context, mongoCollection *mongo.Collection, page int, size int) ([]*models.UserModel, error) {
	option := options.FindOptions{}
	if page == 1 {
		option.SetSkip(0)
		option.SetLimit(int64(size))
	}
	option.SetSkip(int64((page - 1) * size))
	option.SetLimit(int64(size))
	cursor, err := mongoCollection.Find(ctx, bson.D{{}}, &option)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var user models.UserModel
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		userService.Data = append(userService.Data, &user)
	}
	return userService.Data, nil
}

func (userService UserService) GetUserByIdService(ctx context.Context, mongoCollection *mongo.Collection, id string) ([]*models.UserModel, error) {
	filter := bson.M{"_id": id}
	var user models.UserModel
	err := mongoCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	userService.Data = append(userService.Data, &user)
	return userService.Data, nil
}

func (userService UserService) UpdateUserService(ctx context.Context, mongoCollection *mongo.Collection, schema *models.UserModel) ([]*models.UserModel, error) {
	filter := bson.M{"_id": schema.Id}
	_, err := mongoCollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: schema}})
	if err != nil {
		return nil, err
	}
	userService.Data = append(userService.Data, schema)
	return userService.Data, nil
}

func (userService UserService) DeleteUserService(ctx context.Context, mongoCollection *mongo.Collection, id string) (string, error) {
	filter := bson.M{"_id": id}
	_, err := mongoCollection.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}
	return id, nil
}
