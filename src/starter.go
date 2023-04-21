package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"web-hooks/src/serverlogger"
)

func main() {
	start()
}

func start() {
	log.Println("clam-server starting ...")
	initServerBase()
	initServerComponents()
	r := gin.New()
	initGinComponents(r)
	initRouter(r)
	serverHeart(r)
	err := r.Run(":90")
	if err != nil {
		serverlogger.Warn("start clam-server fail", zap.String("err", err.Error()))
		return
	}
	serverlogger.Warn("clam-server started ...")
}

func initServerBase() {
	serverlogger.Init()
}

func initServerComponents() {

}

func initGinComponents(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(serverlogger.LoggerHandler())
	r.Use(Cors())
}

func initRouter(r *gin.Engine) {
	r.POST("/server/update", func(ctx *gin.Context) {
		// doReload
	})
}

func serverHeart(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
