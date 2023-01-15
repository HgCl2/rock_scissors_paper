package handler

import (
	"net/http"

	app "github.com/HgCl2/rock_scissors_paper"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input app.Player

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) signIn(c *gin.Context) {
}
