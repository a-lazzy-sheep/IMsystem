package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags getlist
// @Success 200 {string} GetUserList
// @Router /user/GetUserList [get]
func GetUserList(c *gin.Context) {
	data, err := models.GetUserList()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to get user list",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": data,
	})
}
