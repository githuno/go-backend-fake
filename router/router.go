package router

import (
	api "backend-fake/handlers/api" // 正しくインポートされているか確認

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware()) // CORSミドルウェアの適用

	// JSONファイル操作用のルーティング
	router.GET("/json/:filename", api.GetJsonFile)
	router.PUT("/json/:filename", api.UpdateJsonFile)
	router.POST("/json/:filename", api.CreateJsonFile)

	return router
}
