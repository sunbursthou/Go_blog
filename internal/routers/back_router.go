package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sunbursthou/Go_blog/configs"
)

func InitRouter() {
	// 设置Mode
	gin.SetMode(configs.Cfg.Server.AppMode)

	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"msg": "hello",
			})
		})
	}
	r.Run(configs.Cfg.Server.BackPort)
}
