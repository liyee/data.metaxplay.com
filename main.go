package main

import (
	"encoding/json"
	"time"

	"data.metaxplay.com/common"
	"data.metaxplay.com/help"
	"data.metaxplay.com/initialize"
	"data.metaxplay.com/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.POST("/ob", func(c *gin.Context) {
		data := c.PostForm("data")
		from := c.PostForm("from")
		test := c.PostForm("test")

		if data == "" {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "data is empty",
				"data": "",
			})
			return
		}

		buf := help.Regroup(data)
		ip := c.ClientIP()

		// geoInfo = help.GetClient(ip)
		geoInfo := help.GetClient(ip)

		buf["country_iso_code"] = geoInfo.Country.IsoCode
		buf["country_name_en"] = geoInfo.Country.Names["en"]
		buf["ip_address"] = ip
		buf["test"] = test
		buf["receive_time"] = time.Now().Format("2006-01-02 15:04:05")

		log, _ := json.Marshal(buf)
		dir := common.CONFIG.System.Dir
		help.LogFile(string(log), from, test, dir)

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "succes",
			"data": "",
		})
	})
	r.Run(":" + common.CONFIG.System.Port) // 监听并在 0.0.0.0:8080 上启动服务
}

func init() {
	//初始华配置
	initialize.InitConf()
}
