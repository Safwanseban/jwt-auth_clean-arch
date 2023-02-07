package handlers

import (
	"net/http"

	"github.com/Safwanseban/jwt-auth/pkg/models"
	"github.com/Safwanseban/jwt-auth/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	var user models.User
	logger := utils.GetLogger()

	if err := ctx.ShouldBindJSON(&user); err != nil {
		logger.Error().Err(err).Send()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messege": "bad user-data",
		})
		return
	}
	if err := models.CreateUser(ctx, user); err != nil {
		logger.Error().Err(err).Send()
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": "error encountered",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"messege": "successfully inserted",
	})

}
