package inits

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {

	//配置文件路径
	viper.AddConfigPath("configs")
	//配置文件名称
	viper.SetConfigName("app")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("------------config  app inited------------")
	fmt.Println(json.Marshal(viper.AllSettings()))
	fmt.Println("------------config  app inited------------")
}
