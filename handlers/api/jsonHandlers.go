package handlers

import (
    "net/http"
    "io/ioutil"
    "path/filepath"

    "backend-fake/utils" // ユーティリティ関数を含むパッケージをインポート
    "github.com/gin-gonic/gin" // Ginフレームワークをインポート
)

// GetJsonFile JSONファイルを取得する
func GetJsonFile(c *gin.Context) {
    filename := c.Param("filename")
    filePath := filepath.Join("./json", filename) // ルートディレクトリ(サーバーが実行されているディレクトリ)からの相対パスを指定
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        utils.ResponseError(c, http.StatusNotFound, "File not found")
        return
    }
    c.Data(http.StatusOK, "application/json; charset=utf-8", content)
}
// UpdateJsonFile JSONファイルを更新する
func UpdateJsonFile(c *gin.Context) {
    filename := c.Param("filename")
    filePath := filepath.Join("./json", filename) // ルートディレクトリ(サーバーが実行されているディレクトリ)からの相対パスを指定
    bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
    defer c.Request.Body.Close()
    err := ioutil.WriteFile(filePath, bodyBytes, 0644)
    if err!= nil {
        utils.ResponseError(c, http.StatusInternalServerError, "Failed to update file")
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "File updated successfully"})
}

// CreateJsonFile JSONファイルを作成する
func CreateJsonFile(c *gin.Context) {
    filename := c.Param("filename")
    filePath := filepath.Join("./json", filename)
    bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
    defer c.Request.Body.Close()
    err := ioutil.WriteFile(filePath, bodyBytes, 0644)
    if err!= nil {
        utils.ResponseError(c, http.StatusInternalServerError, "Failed to create file")
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "File created successfully"})
}
