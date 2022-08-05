package main

import (
	"encoding/json"
	"time"

	"data.metaxplay.com/help"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
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
		buf["geo"] = help.GetClient(ip)
		buf["test"] = test
		buf["receive_time"] = time.Now().Format("2006-01-02 15:04:05")

		log, _ := json.Marshal(buf)
		help.LogFile(string(log), from, test)

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "succes",
			"data": "",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func init() {
	//初始华配置
	initialize.InitConf()
}
