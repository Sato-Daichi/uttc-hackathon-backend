package controller

import (
	"app/usecase"
	"encoding/json"
	"log"
	"net/http"
)

// チャンネルidから全メッセージを取得
func GetMessagesByChannelID(w http.ResponseWriter, r *http.Request) {
	// チャンネルidを取得
	channelID := r.URL.Query().Get("channel")

	// メッセージを取得
	messages, err := usecase.GetMessagesByChannelID(channelID)
	if err != nil {
		log.Printf("fail: usecase.GetAllMessagesByChannelID, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// JSONに変換
	jsonBytes, err := json.Marshal(messages)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// レスポンスに書き込み
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonBytes)
}
