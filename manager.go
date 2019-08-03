package bullet_text_chatroom

import "log"

const (
	timeLayout      = "2006-01-02 15:04:05"
	broadcastRoomId = "GodSeeUTomorrow"
)

type Msg struct {
	Data   string `json:"data"`
	From   string `json:"from"`
	RoomId string `json:"roomId"`
	Time   string `json:"time"`
}

//ClientManager 客户端管理器
type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan Msg
	register   chan *Client
	unregister chan *Client
}

// 初始化manager全局变量
var manager = ClientManager{
	clients:    make(map[*Client]bool),
	broadcast:  make(chan Msg),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

func (manager *ClientManager) start() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[System] error:%+v", err)
		}
	}()

	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.sendChan)
				_ = conn.socket.Close()
				delete(manager.clients, conn)
			}
		case msg := <-manager.broadcast:
			for client := range manager.clients {
				if broadcastRoomId == msg.RoomId || client.roomId == msg.RoomId {
					client.sendChan <- msg
				}
			}
		}
	}
}

func (manager *ClientManager) Join(client *Client) {
	manager.register <- client
}

func (manager *ClientManager) Leave(client *Client) {
	manager.unregister <- client
}

func (manager *ClientManager) DealMsg(msg Msg) {
	manager.broadcast <- msg
}
