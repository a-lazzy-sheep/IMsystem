package models

import (
	"encoding/json"
	"fmt"
	// "log"
	// "net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
)

//消息
type Message struct {
	gorm.Model
	UserId     int64  //发送者
	TargetId   int64  //接受者
	Type       int    //发送类型  1私聊  2群聊  3心跳
	Media      int    //消息类型  1文字 2表情包 3语音 4图片 /表情包
	Content    string //消息内容
	CreateTime uint64 //创建时间
	ReadTime   uint64 //读取时间
	Pic        string
	Url        string
	Desc       string
	Amount     int //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

// const (
// 	HeartbeatMaxTime = 1 * 60
// )

type Node struct {
	Conn          *websocket.Conn //连接
	Addr          string          //客户端地址
	FirstTime     uint64          //首次连接时间
	HeartbeatTime uint64          //心跳时间
	LoginTime     uint64          //登录时间
	DataQueue     chan []byte     //消息
	GroupSets     set.Interface   //好友 / 群
}

//映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

//读写锁
var rwLocker sync.RWMutex

//	需要 ：发送者ID ，接受者ID ，消息类型，发送的内容，发送类型
func Chat(writer http.ResponseWriter, request *http.Request) {
	//1.  获取参数 并 检验 token 等合法性
	//token := query.Get("token")
	query := request.URL.Query()
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	//msgType := query.Get("type")
	//targetId := query.Get("targetId")
	//	context := query.Get("context")
	isvalida := true //checkToke()  待.........
	conn, err := (&websocket.Upgrader{
		//token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//2.获取conn
	currentTime := uint64(time.Now().Unix())
	node := &Node{
		Conn:          conn,
		Addr:          conn.RemoteAddr().String(), //客户端地址
		HeartbeatTime: currentTime,                //心跳时间
		LoginTime:     currentTime,                //登录时间
		DataQueue:     make(chan []byte, 50),
		GroupSets:     set.New(set.ThreadSafe),
	}
	//3. 用户关系
	//4. userid 跟 node绑定 并加锁
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()
	//5.完成发送逻辑
	go sendProc(node)
	//6.完成接受逻辑
	go recvProc(node)

	sendMsg(userId, []byte("欢迎进入聊天系统"))

}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			fmt.Println("[ws]sendProc >>>> msg :", string(data))
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		// 从WebSocket连接中读取消息
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("[ws] recvProc <<<<< ", string(data))
		dispatch(data)
		// broadMsg(data) //todo 将消息广播到局域网
	}
}

// var tcpsendChan chan []byte = make(chan []byte, 1024)

// func broadMsg(data []byte) {
// 	tcpsendChan <- data
// }

// func init() {
// 	go tcpSendProc()
// 	go tcpRecvProc()
// 	fmt.Println("init goroutine ")
// }

// //完成tcp数据发送协程
// func tcpSendProc() {
// 	conn, err := net.Dial("tcp", "localhost:8080")
// 	if err != nil {
// 		log.Printf("tcpSendProc err: %v", err)
// 	}
// 	defer conn.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for {
// 		select {
// 		case data := <-tcpsendChan:
// 			fmt.Println("tcpSendProc  data :", string(data))
// 			_, err := conn.Write(data)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 		}
// 	}

// }

// //完成tcp数据接收协程
// func tcpRecvProc() {
// 	listener, err := net.Listen("tcp", "localhost:8080")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer listener.Close()
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		go handleTCPConnection(conn)
// 	}
// }

// func handleTCPConnection(conn net.Conn) {
// 	defer conn.Close()
// 	for {
// 		var buf [512]byte
// 		n, err := conn.Read(buf[0:])
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Println("tcpRecvProc  data :", string(buf[0:n]))
// 		dispatch(buf[0:n])
// 	}
// }

//后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	msg.CreateTime = uint64(time.Now().Unix())
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: //私信
		fmt.Println("dispatch  data :", string(data))
		sendMsg(msg.TargetId, data)
	case 2: //群发
		sendGroupMsg(msg.TargetId, data) //发送的群ID ，消息内容
		// case 4: // 心跳
		// 	node.Heartbeat()
		//case 4:
		//
	}
}


func sendMsg(userId int64, msg []byte) {

	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	} else {
		fmt.Println("sendMsg  userId not found")
	}	
}

func sendGroupMsg(targetId int64, msg []byte) {
	fmt.Println("开始群发消息")
	userIds := SearchUserByGroupId(uint(targetId))
	for i := 0; i < len(userIds); i++ {
		//排除给自己的
		if targetId != int64(userIds[i]) {
			sendMsg(int64(userIds[i]), msg)
		}

	}
}