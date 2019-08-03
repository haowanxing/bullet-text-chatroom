package bullet_text_chatroom

import (
	"errors"
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
	"time"
)

const TemplateDir = "template/"
const PageHome = "index.html"

func Serve(port int) {
	if port < 0 || port > 65535 {
		log.Println(errors.New("端口号不正确"))
		return
	}
	// 开启socket管理器
	go manager.start()
	go server(port)
}

func server(port int) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[System] Server error: %+v", err)
		}
	}()

	// 绑定WebSocket服务路径
	http.Handle("/ws", websocket.Handler(WsHandler))
	// Web节点
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles(TemplateDir + PageHome)
		if err != nil {
			log.Printf("[System] Server error: %+v", err)
		}
		_ = t.Execute(writer, nil)
	})
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func WsHandler(ws *websocket.Conn) {
	client := &Client{
		id:       time.Now().String(),
		roomId:   "",
		socket:   ws,
		sendChan: make(chan Msg, 1024),
	}
	manager.register <- client
	go client.shit()
	client.eat()
}
