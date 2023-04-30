package controllers

import (
	"app-user/models"
	"app-user/services"
	"context"
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
// @Success 200 {object} object{code="200",status="OK",data=[]models.UserModel,"executon_time"=0.0} "OK"
// @Router /api/v1/user [post]
func (userController UserController) CreateUserController(c *gin.Context) {
	start := time.Now()
	user := models.UserModel{}
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"code": "400", "status": "FAILED", "message": "parsing body error"})
		return
	}
	datas, err := userController.UserService.CreateUserService(context.TODO(), userController.MongoCollection, user)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"code": "400", "status": "FAILED", "message": "error on insert"})
		return
	}
	c.JSON(200, gin.H{"code": "200", "status": "OK", "data": datas, "execution_time": time.Since(start).Seconds()})
}

func (userController UserController) GetUserController() {}

func (userController UserController) GetUserByIdController() {}

func (userController UserController) UpdateUserController() {}

func (userController UserController) DeleteUserController() {}
