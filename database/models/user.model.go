package models


type UserLogin struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}


type User struct {
	Username string `json:"username" bson:"username"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Avatar string `json:"avatar" bson:"avatar"`
}