package main

import (
	bullettextchatroom "bullet-text-chatroom"
)

func main() {
	bullettextchatroom.Serve(4000)
	select {}
	//time.Sleep(time.Second*3600)
}
