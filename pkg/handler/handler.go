package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	hall := router.Group("/hall")
	{
		hall.GET("/", h.hall)
		hall.POST("/create-room", h.createRoom)
		hall.GET("/gameroom", h.getGameRoom)
		hall.POST("/gameroom", h.postGameRoom)
	}

	return router
}
