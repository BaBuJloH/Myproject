package server

import (
	"net/http"
	"MyProject/service"


	"github.com/gin-gonic/gin"
)

type Server struct {
	Port string
}

func New(port string, srvc *service.Service) *gin.Engine {
	g := gin.Default()

	g.Get("/about", func(ctx *gin.Context) {
		ctx.JSON(http.StatusNoContent, nil)
	})

	{
		users := user.Handlers{*srvc}
		g.GET("/users", users.List)
		g.GET("/user/:user_id", users.Get)
		g.POST("/user", users.Add)
		g.DELETE("/user/:user_id", users.Del)
	}
	return g
}
