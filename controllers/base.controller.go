package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security 	BasicAuth
// @Summary 	Information about the method
// @Tags 		GroupA
// @Accept 		json
// @Produce 	json
// @Param 		name query string true "enter a name"
// @Success 	200 {object} dto.SuccessResponse
// @Failure 	400 {object} dto.ErrorResponse
// @Failure 	401 {object} dto.ErrorResponse
// @Router 		/msg [get]
func MSGController(c *gin.Context) {

	// qiymatni o'qish
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{"Msg": "Hello " + name})
}
