package test

import (
	"NewGateServer/WeHub"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func BenchmarkBaseMessage(b *testing.B) {
	var d = WeHub.BaseMessage{}
	for i := 0; i < b.N; i++ {
		bm, err := d.MarshalJSON()
		if err != nil {
			b.Log(err)
		} else {
			var temp WeHub.BaseMessage
			err := temp.UnmarshalJSON(bm)
			if err != nil {
				b.Log(err)
			}
		}

	}
}

func TestCloseChan(t *testing.T) {
	var c = make(chan bool)
	go func() {
		time.Sleep(time.Second * 1)
		close(c)
	}()

	go func() { c <- true }()

	for {
		select {
		case res, ok := <-c:
			if ok {
				fmt.Println("收到消息", strconv.FormatBool(res))
			} else {
				fmt.Println("通道关闭")
				return
			}
		}
	}
}
