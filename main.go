package main

import (
    "backend-fake/router"
)

func main() {
    router := router.SetupRouter()
    router.Run(":8080") // サーバーのポート番号を変更可能
}
