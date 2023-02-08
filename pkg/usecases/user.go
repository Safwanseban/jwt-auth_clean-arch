package usecases

import (
	"github.com/Safwanseban/jwt-auth/pkg/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context, user models.User) error {

	userpas, err := user.HashUserPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(userpas)
	return user.InsertOne()
}

func GetUserData(ctx *gin.Context, user models.User) error {

	userdata, err := user.FindUser()
	if err != nil {
		return err
	}

	if err := user.CheckPassword(userdata.Password); err != nil {
		return err
	}
	return nil
}
