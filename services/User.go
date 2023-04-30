package services

import (
	"app-user/models"
	"app-user/utils"
	"context"
	"crypto/md5"
	"encoding/hex"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	Data []models.UserModel
}

func InitUserService() *UserService {
	return &UserService{}
}

func (userService UserService) CreateUserService(ctx context.Context, mongoCollection *mongo.Collection, schema models.UserModel) ([]models.UserModel, error) {
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
	userService.Data = append(userService.Data, schema)
	return userService.Data, nil
}

func (userService UserService) GetUser(mongoCollection *mongo.Collection, page int, size int) ([]models.UserModel, error) {
	return userService.Data, nil
}

func (userService UserService) GetUserByIdService(mongoCollection *mongo.Collection, id string) ([]models.UserModel, error) {
	return userService.Data, nil
}

func (userService UserService) UpdateUserService(mongoCollection *mongo.Collection, id string, schema *models.UserModel) ([]models.UserModel, error) {
	return userService.Data, nil
}

func (userService UserService) DeleteUserService(mongoCollection *mongo.Collection, id string) ([]models.UserModel, error) {
	return userService.Data, nil
}
