package main

import (
	"child-coding-platform/backend/api"
	"child-coding-platform/backend/database"
	"fmt"
	"net/http"
)

func main() {

	// 首先连接数据库
	database.Connect()

	// 然后设置路由
	http.HandleFunc("/register", api.RegisterUser)
	http.HandleFunc("/login", api.LoginUser)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "欢迎来到少儿编程平台!")
	})

	// 打印启动服务器的消息
	fmt.Println("服务器启动在 http://localhost:8081")

	// 最后，启动服务器（这将会阻塞，直到服务器停止）
	http.ListenAndServe(":8081", nil)

}
