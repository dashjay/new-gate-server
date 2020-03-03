package test

import (
	"NewGateServer/WeHub"
	"testing"
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
