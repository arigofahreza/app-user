package controllers

import (
	"app-user/models"
	"app-user/services"
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	MongoCollection *mongo.Collection
	UserService     *services.UserService
}

func InitUserController(mongoCollection *mongo.Collection) *UserController {
	return &UserController{
		MongoCollection: mongoCollection,
		UserService:     services.InitUserService(),
	}
}

// @Summary Add User
// @Description create new user
// @Param body body models.UserModel true "body"
// @Tags User
// @Accept  json
// @Produce  json
// @Router /api/v1/user [post]
func (userController UserController) CreateUserController(c *gin.Context) {
	start := time.Now()
	user := models.UserModel{}
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"code": 400, "status": "FAILED", "message": "parsing body error", "execution_time": time.Since(start).Seconds()})
		return
	}
	datas, err := userController.UserService.CreateUserService(context.TODO(), userController.MongoCollection, user)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"code": 500, "status": "FAILED", "message": "error on insert", "execution_time": time.Since(start).Seconds()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "status": "OK", "data": datas, "execution_time": time.Since(start).Seconds()})
}

// @Summary Get User
// @Description get all user
// @Param   page     query    string     true        "1"
// @Param   size     query    string     true        "100"
// @Tags User
// @Accept  json
// @Produce  json
// @Router /api/v1/user [get]
func (userController UserController) GetUserController(c *gin.Context) {
	start := time.Now()
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"code": 400, "status": "FAILED", "message": "page required", "execution_time": time.Since(start).Seconds()})
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"code": 400, "status": "FAILED", "message": "size required", "execution_time": time.Since(start).Seconds()})
		return
	}
	datas, err := userController.UserService.GetUserService(context.TODO(), userController.MongoCollection, page, size)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"code": 500, "status": "FAILED", "message": "error on retrieving data", "execution_time": time.Since(start).Seconds()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "status": "OK", "data": datas, "execution_time": time.Since(start).Seconds()})
}

// @Summary Get User by id
// @Description get user by id
// @Param   id     path    string     true        "id"
// @Tags User
// @Accept  json
// @Produce  json
// @Router /api/v1/user/{id} [get]
func (userController UserController) GetUserByIdController(c *gin.Context) {
	start := time.Now()
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"code": 400, "status": "FAILED", "message": "id not found", "execution_time": time.Since(start).Seconds()})
		return
	}
	datas, err := userController.UserService.GetUserByIdService(context.TODO(), userController.MongoCollection, id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"code": 500, "status": "FAILED", "message": "error on retrieving data", "execution_time": time.Since(start).Seconds()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "status": "OK", "data": datas, "execution_time": time.Since(start).Seconds()})
}

// @Summary Update User
// @Description update existing user
// @Param body body models.UserModel true "body"
// @Tags User
// @Accept  json
// @Produce  json
// @Router /api/v1/user [put]
func (userController UserController) UpdateUserController(c *gin.Context) {
	start := time.Now()
	var schema models.UserModel
	err := c.BindJSON(&schema)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"code": 400, "status": "FAILED", "message": "error on parsing body", "execution_time": time.Since(start).Seconds()})
		return
	}
	datas, err := userController.UserService.UpdateUserService(context.TODO(), userController.MongoCollection, &schema)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"code": 500, "status": "FAILED", "message": "error on updating data", "execution_time": time.Since(start).Seconds()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "status": "OK", "data": datas, "execution_time": time.Since(start).Seconds()})
}

// @Summary Delete User
// @Description delete user by id
// @Param   id     path    string     true        "id"
// @Tags User
// @Accept  json
// @Produce  json
// @Router /api/v1/user/{id} [delete]
func (userController UserController) DeleteUserController(c *gin.Context) {
	start := time.Now()
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"code": 400, "status": "FAILED", "message": "id not found", "execution_time": time.Since(start).Seconds()})
		return
	}
	datas, err := userController.UserService.DeleteUserService(context.TODO(), userController.MongoCollection, id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"code": 500, "status": "FAILED", "message": "error on retrieving data", "execution_time": time.Since(start).Seconds()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "status": "OK", "data": datas, "execution_time": time.Since(start).Seconds()})
}
