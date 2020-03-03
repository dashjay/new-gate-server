package WeHub

import (
	"github.com/silenceper/wechat"
)

type Servers struct {
	Wc           *wechat.Wechat
	cfg          *Config
	ResponseTime []float64
}
