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
	http.HandleFunc("/messages", controller.GetMessagesByChannelId)

	// user_idからworkspacesを取得
	http.HandleFunc("/workspaces", controller.GetWorkspacesByUserId)

	// workspace_idからchannelsを取得
	http.HandleFunc("/channels", controller.GetChannelsByWorkspaceId)

	// message_idからメッセージを削除
	http.HandleFunc("/message/delete", controller.DeleteMessage)

	// messageを投稿
	http.HandleFunc("/message/post", controller.PostMessage)

	// messageをupdate
	http.HandleFunc("/message/update", controller.UpdateMessage)

	// login時にusernameを取得
	http.HandleFunc("/login", controller.UserLogin)

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
