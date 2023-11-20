package user

import (
	"MyProject/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct{
	service.Service
}

//Add - POST/user - create user
func (h Handlers) Add(ctx *gin.Context){
	usr := user.User{
		if err:= ctx.BindJSON(&usr); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := h.Repositories.Users.Insert(usr); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"full_path": ctx.FullPath(),
			"method": ctx.Request.Method,
		})
	}
}