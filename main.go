/**
* @Auth:ShenZ
* @Description:
* @CreateDate:2022/06/15 10:57:44
 */
package main

import (
	"github.com/spf13/viper"
	"mychat/clients"
	"mychat/inits"
	"mychat/models"
	"mychat/router"
)

func main() {

	inits.InitConfig()
	inits.InitTimer()
	clients.InitMySQL()
	clients.InitRedis()
	models.InitUdpProc()

	r := router.Router()
	r.Run(viper.GetString("port.server"))

}
