package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (h *Handler) googleLogin(c *gin.Context) {
	defer c.Request.Body.Close()

	// parse the GoogleJWT that was POSTed from the front-end
	type parameters struct {
		GoogleJWT *string
	}
	decoder := json.NewDecoder(c.Request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		newErrorResponse(c, 500, "Couldn't decode parameters")
		return
	}

	// Validate the JWT is valid
	claims, err := h.services.ValidateGoogleJWT(*params.GoogleJWT)
	if err != nil {
		newErrorResponse(c, 403, "Invalid google auth")
		return
	}

	// create a JWT for OUR app and give it back to the client for future requests
	tokenString, err := h.services.GenerateToken(claims.Email)
	if err != nil {
		newErrorResponse(c, 500, "Couldn't make authentication token")
		return
	}

	c.JSON(200, struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	})
}
