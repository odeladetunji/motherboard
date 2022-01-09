package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

// Easier to get running with CORS. Thanks for help @Vindexus and @erkie
var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func main() {
	router := gin.Default()
	// router.LoadHTMLGlob("../templates/*")
	router.LoadHTMLFiles("../templates/chat.html", "../templates/iniciatecall.html",
	 "../templates/chat1.html", "../templates/login.html", "../templates/videocall.html")

	router.GET("/chat", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.html", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/chat1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat1.html", gin.H{
			"title": "Main website2",
		})
	})

	router.GET("/videocall", func(c *gin.Context) {
		c.HTML(http.StatusOK, "videocall.html", gin.H{
			"title": "MotherBoard",
		})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "MotherBoard",
		})
	})

	router.GET("/iniciatecall", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iniciatecall.html", gin.H{
			"title": "MotherBoard",
		})
	})

	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		// s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "connect-to-chat-server", func(s socketio.Conn, msg map[string]int) string {
		fmt.Println("connect-to-chat-server:", msg["sender_id"])
		// s.Emit("reply", "have "+msg)
		return "connect-to-chat-server"
	})

	// server.OnError("/", func(s socketio.Conn, e error) {
	// 	fmt.Println("meet error:", e)
	// })

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	defer server.Close()

	router.GET("/socket.io/*any", gin.WrapH(server))

	// if err := router.Run(":8000"); err != nil {
	// 	log.Fatal("failed run app: ", err)
	// }

	router.Run(":9001")
}
