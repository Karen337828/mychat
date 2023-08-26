package inits

import (
	"github.com/spf13/viper"
	"mychat/models"
	"time"
)

type TimerFunc func(interface{}) bool

func InitTimer() {

	delay := time.Duration(viper.GetInt("timeout.DelayHeartbeat")) * time.Second
	tick := time.Duration(viper.GetInt("timeout.HeartbeatHz")) * time.Second
	var cleanConnectFunc TimerFunc = models.CleanConnection
	Timer(delay, tick, cleanConnectFunc, "")
}

// Timer /**
func Timer(delay time.Duration, tick time.Duration, fun TimerFunc, param interface{}) {
	go func() {
		if fun == nil {
			return
		}
		t := time.NewTimer(delay)
		for {
			select {
			case <-t.C:
				if fun(param) == false {
					return
				}
				t.Reset(tick)
			}
		}
	}()
}
