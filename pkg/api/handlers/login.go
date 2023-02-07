package handlers

import (
	"net/http"

	"github.com/Safwanseban/jwt-auth/pkg/models"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user models.User

	if err:=ctx.ShouldBindJSON(&user); err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"messege":"error binding json",
		})
		return
	}
	if err:=models.CreateUser(ctx,user);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"err":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"messege":"successfully inserted",
	})

}
