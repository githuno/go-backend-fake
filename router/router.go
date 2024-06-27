package router

import (
    "github.com/gin-gonic/gin"
    api "backend-fake/handlers/api" // 正しくインポートされているか確認
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // JSONファイル操作用のルーティング
    router.GET("/json/:filename", api.GetJsonFile)
    router.PUT("/json/:filename", api.UpdateJsonFile)
    router.POST("/json/:filename", api.CreateJsonFile)

    return router
}
