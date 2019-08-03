package bullet_text_chatroom

import (
	"testing"
	"time"
)

func TestServe(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	Serve(8080)
	time.Sleep(time.Second * 5)
	t.Log("5秒后打印此日志，表示一切Ok")
}
