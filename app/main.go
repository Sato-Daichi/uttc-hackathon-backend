package main

import (
	"app/controller"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// 最初に呼ばれる関数
func init() {}

func main() {
	http.HandleFunc("/signup", controller.UserSignUp)

	// 全ユーザーを取得
	http.HandleFunc("/users/all", controller.GetAllUsers)

	// チャンネルidから全メッセージを取得
	http.HandleFunc("/messages", controller.GetMessagesByChannelID)

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
