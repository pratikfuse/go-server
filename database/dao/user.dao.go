package dao

import (
	. "fiberapp/database"
	"fiberapp/database/models"
)

const Collection = "Users"

func Authenticate(user models.UserLogin) (models.User, error) {
	var userData models.User
	err := Conn.Db.C(Collection).Find(user).One(&userData)
	return userData, err
}
