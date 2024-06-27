package utils

import (
    "github.com/gin-gonic/gin"
)

// ResponseError HTTPステータスコードとエラーメッセージをレスポンスとして返す
func ResponseError(c *gin.Context, statusCode int, message string) {
    c.JSON(statusCode, gin.H{"error": message})
}
