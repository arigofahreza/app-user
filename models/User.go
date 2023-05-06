package models

type UserModel struct {
	Id          string `json:"_id" bson:"_id"`
	Username    string `json:"username" bson:"username"`
	Name        string `json:"name" bson:"name"`
	DateOfBirth int    `json:"date_of_birth" bson:"date_of_birth"`
	Email       string `json:"email" bson:"email"`
	Password    string `json:"password" bson:"password"`
	IsVerified  bool   `json:"is_verified" bson:"is_verified"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
}

type ResponseModel struct {
	Code          int
	Status        string
	Data          interface{}
	ExecutionTime float64
}
