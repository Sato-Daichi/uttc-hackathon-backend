package controller

import (
	"app/usecase"
	"encoding/json"
	"log"
	"net/http"
)

// チャンネルidから全メッセージを取得
func GetMessagesByChannelId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// チャンネルidを取得
		channelId := r.URL.Query().Get("channel")
		if channelId == "" {
			log.Printf("fail: r.URL.Query().Get, %v\n", channelId)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// メッセージを取得
		messages, err := usecase.GetMessagesByChannelId(channelId)
		if err != nil {
			log.Printf("fail: usecase.GetAllMessagesByChannelId, %v\n", err)
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
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
