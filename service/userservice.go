package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
// @param email query string true "Email of the user"
// @param name query string true "Name of the user"
// @param password query string true "Password of the user"
// @Success 200 {string} Register
// @Router /user/Register [post]
func Register(c *gin.Context) {
	var user models.UserBasic
	user.Email = c.Request.FormValue("email")
	user.Name = c.Request.FormValue("name")
	user.Password = c.Request.FormValue("password")
	// rePassword := c.Request.FormValue("rePassword")
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
	// if user.Password != rePassword {
	// 	log.Printf("Passwords do not match and user.Password is %s and rePassword is %s", user.Password, rePassword)
	// 	c.JSON(400, gin.H{
	// 		"error": "Passwords do not match",
	// 	})
	// 	return
	// }
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
		"code" : 0,
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
	if id64, err = strconv.ParseUint(c.Request.FormValue("id"), 10, 64); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	user.ID = uint(id64)

	user.Name = c.Request.FormValue("name")
	user.Email = c.Request.FormValue("email")
	user.Password = c.Request.FormValue("password")
	user.Phone = c.Request.FormValue("phone")
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
	user.Email = c.Request.FormValue("email")
	user.Password = c.Request.FormValue("password")
	if u, err := models.FindUserByEmailAndPassword(user.Email, user.Password); err != nil {
		c.JSON(401, gin.H{
			"error": "Invalid email or password",
		})
		return
	} else {
		// 生成token
		token, err:= utils.GenerateToken(user.Email, user.Password)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to generate token",
			})
			return
		}
		c.JSON(200, gin.H{
			"code": "0",
			"message": "Login successfully",
			"token" :  token,
			"data": u,
		})
	}
}


// 防止跨域站点伪造请求
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
    WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
func SendMsg(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(c, ws)
}

func WebsocketHandler(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()
	for {
		// 读取消息
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("Received: %s", message)

		// 发送消息
		if err := ws.WriteMessage(messageType, message); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}

// SearchFriends
// @Summary search friends
// @Tags searchfriends
// @param userId query string true "Id of the user"
// @Success 200 {string} Register
// @Router /searchFriends [post]
func SearchFriends(c *gin.Context) {
	id, _ := strconv.Atoi(c.Request.FormValue("userId"))
	users := models.SearchFriend(uint(id))
	// c.JSON(200, gin.H{
	// 	"code":    0, //  0成功   -1失败
	// 	"message": "查询好友列表成功！",
	// 	"data":    users,
	// })
	utils.RespOKList(c.Writer, users, len(users))
}

// AddFriend
// @Summary add friends
// @Tags addfriends
// @param userId query string true "Id of the user"
// @param targetName query string true "targetName of the user"
// @Success 200 {string} Register
// @Router /contact/addfriend [post]
func AddFriend(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))
	targetName := c.Request.FormValue("targetName")
	//targetId, _ := strconv.Atoi(c.Request.FormValue("targetId"))
	code, msg := models.AddFriend(uint(userId), targetName)
	if code == 0 {
		utils.RespOK(c.Writer, code, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}

//加载群列表
func LoadCommunity(c *gin.Context) {
	ownerId, _ := strconv.Atoi(c.Request.FormValue("ownerId"))
	//	name := c.Request.FormValue("name")
	data, msg := models.LoadCommunity(uint(ownerId))
	if len(data) != 0 {
		utils.RespList(c.Writer, 0, data, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}

//新建群
func CreateCommunity(c *gin.Context) {
	ownerId, _ := strconv.Atoi(c.Request.FormValue("ownerId"))
	name := c.Request.FormValue("name")
	icon := c.Request.FormValue("icon")
	desc := c.Request.FormValue("desc")
	community := models.Community{}
	community.OwnerId = uint(ownerId)
	community.Name = name
	community.Img = icon
	community.Desc = desc
	code, msg := models.CreateCommunity(community)
	if code == 0 {
		utils.RespOK(c.Writer, code, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}

//加入群 userId uint, comId uint
func JoinCommunity(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))
	comId := c.Request.FormValue("comId")

	//	name := c.Request.FormValue("name")
	data, msg := models.JoinCommunity(uint(userId), comId)
	if data == 0 {
		utils.RespOK(c.Writer, data, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}

func FindByID(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))

	//	name := c.Request.FormValue("name")
	data := models.FindByID(uint(userId))
	utils.RespOK(c.Writer, data, "ok")
}

func RedisMsg(c *gin.Context) {
	userIdA, _ := strconv.Atoi(c.PostForm("userIdA"))
	userIdB, _ := strconv.Atoi(c.PostForm("userIdB"))
	start, _ := strconv.Atoi(c.PostForm("start"))
	end, _ := strconv.Atoi(c.PostForm("end"))
	isRev, _ := strconv.ParseBool(c.PostForm("isRev"))
	res := models.RedisMsg(int64(userIdA), int64(userIdB), int64(start), int64(end), isRev)
	utils.RespOKList(c.Writer, "ok", res)
}

func MsgHandler(c *gin.Context, ws *websocket.Conn) {
	for {
		msg, err := utils.Subscribe(c, utils.PublicKey)
		if err != nil {
			fmt.Println(" MsgHandler 发送失败", err)
		}

		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			log.Fatalln(err)
		}
	}
}