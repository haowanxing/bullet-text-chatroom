package bullet_text_chatroom

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"time"
)

const (
	clientEventName = "name"
	clientEventRoom = "room"
	clientEventText = "text"
)

type Client struct {
	id       string
	roomId   string
	socket   *websocket.Conn
	sendChan chan Msg
}

type ClientMsg struct {
	Event string `json:"event"`
	Time  string `json:"time"`
	Data  string `json:"data"`
}

func (client *Client) eat() {
	for msg := range client.sendChan {
		msgByte, _ := json.Marshal(msg)

		if _, err := client.socket.Write(msgByte); err != nil {
			log.Printf("[System] write msg err: %+v", err)
			manager.Leave(client)
			break
		}
	}
	log.Printf("[System] client's websocket closed")
}

func (client *Client) shit() {
	for {
		var shit string
		if err := websocket.Message.Receive(client.socket, &shit); err != nil {
			if err != io.EOF {
				log.Printf("[System] can not catch client's shit. so dirty: %+v", err)
			} else {
				log.Printf("[System] client Leave: %+v", client.id)
			}
			manager.Leave(client)
			break
		}
		var msg = &ClientMsg{Time: time.Now().Format(timeLayout)}
		if err := json.Unmarshal([]byte(shit), msg); err != nil {
			log.Printf("[System] can not json client's shit. so bad: %+v", err)
			manager.Leave(client)
			break
		}

		switch msg.Event {
		case clientEventName:
			if 0 != len(msg.Data) {
				client.id = msg.Data
			}
		case clientEventRoom:
			if broadcastRoomId != msg.Data && client.roomId != msg.Data {
				client.roomId = msg.Data
				manager.DealMsg(Msg{
					Event:  systemEventNotice,
					Data:   client.id + "进入聊天室",
					From:   "系统消息",
					RoomId: client.roomId,
					Time:   time.Now().Format(timeLayout),
				})
			}
		case clientEventText:
			if 0 != len(msg.Data) {
				manager.DealMsg(Msg{
					Event:  msg.Event,
					Data:   msg.Data,
					From:   client.id,
					RoomId: client.roomId,
					Time:   time.Now().Format(timeLayout),
				})
			}
		}
	}
}
