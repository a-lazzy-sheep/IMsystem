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

// Register
// @Summary Register a new user
// @Tags register
// @param name query string true "Name of the user"
// @param email query string true "Email of the user"
// @param password query string true "Password of the user"
// @Success 200 {string} Register
// @Router /user/Register [post]
func Register(c *gin.Context) {
	var user models.UserBasic
	user.Name = c.Query("name")
	user.Email = c.Query("email")
	user.Password = c.Query("password")
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}
	if err := models.CreateUser(&user); err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User created successfully",
	})
}

// DeleteUser
// @Summary Delete a user
// @Tags DeleteUser
// @param name query string true "Name of the user"
// @Success 200 {string} DeleteUser
// @Router /user/DeleteUser [delete]
func DeleteUser(c *gin.Context) {
	var user models.UserBasic
	user.Name = c.Query("name")
	if err := models.DeleteUser(&user); err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to delete user",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User deleted successfully",
	})
}