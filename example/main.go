package main

import chatroom "github.com/haowanxing/bullet-text-chatroom"

func main() {
	chatroom.Serve(4000)
	select {}
	//time.Sleep(time.Second*3600)
}
