package service

import (
	"ginchat/models"
	"log"
	"strconv"

	"github.com/asaskevich/govalidator"
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
// @param rePassword query string true "Re-enter password of the user"
// @Success 200 {string} Register
// @Router /user/Register [post]
func Register(c *gin.Context) {
	var user models.UserBasic
	user.Name = c.Query("name")
	user.Email = c.Query("email")
	user.Password = c.Query("password")
	rePassword := c.Query("rePassword")
	if _, err := models.FindUserByEmail(user.Email); err == nil {
		c.JSON(400, gin.H{
			"error": "Email already exists",
		})
		return
	}
	if _, err := models.FindUserByName(user.Name); err == nil {
		c.JSON(400, gin.H{
			"error": "Username already exists",
		})
		return
	}
	if user.Password != rePassword {
		c.JSON(400, gin.H{
			"error": "Passwords do not match",
		})
		return
	}
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
// UpdateUser
// @Summary Update a user
// @Tags UpdateUser
// @param id query string true "ID of the user"
// @param name query string true "Name of the user"
// @param email query string true "Email of the user"
// @param password query string true "Password of the user"
// @param phone query string true "Phone of the user"
// @Success 200 {string} UpdateUser
// @Router /user/UpdateUser [put]
func UpdateUser(c *gin.Context) {
	var user models.UserBasic
	var err error

	// 将字符串类型的id转换为uint类型
	var id64 uint64
	if id64, err = strconv.ParseUint(c.Query("id"), 10, 64); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	user.ID = uint(id64)

	user.Name = c.Query("name")
	user.Email = c.Query("email")
	user.Password = c.Query("password")
	user.Phone = c.Query("phone")
	result, err := govalidator.ValidateStruct(&user)
	if err != nil {
		log.Printf("Invalid data format: %v", err)
		c.JSON(400, gin.H{
			"error": "Invalid data format", 
			"result" : result,
		})
		return
	}
	if err := models.UpdateUser(&user); err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update user because of wrong id or columns",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User updated successfully",
	})
}



// Login
// @Summary User login
// @Tags Login
// @param email query string true "Email of the user"
// @param password query string true "Password of the user"
// @Success 200 {string} Login
// @Router /user/Login [post]
func Login(c *gin.Context) {
	var user models.UserBasic
	user.Email = c.Query("email")
	user.Password = c.Query("password")
	if u, err := models.FindUserByEmailAndPassword(user.Email, user.Password); err != nil {
		c.JSON(401, gin.H{
			"error": "Invalid email or password",
		})
		return
	} else {
		// 生成token
		// token, err:= models.GenerateToken()
		token := "dfsajklvhc"
		// if err != nil {
		// 	c.JSON(500, gin.H{
		// 		"error": "Failed to generate token",
		// 	})
		// 	return
		// }
		c.JSON(200, gin.H{
			"token" :  token,
			"message": "Login successful",
			"User information": u,
		})
	}
}