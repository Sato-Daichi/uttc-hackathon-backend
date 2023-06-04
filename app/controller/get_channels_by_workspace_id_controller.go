package controller

import (
	"app/usecase"
	"encoding/json"
	"log"
	"net/http"
)

// ワークスペースidからチャンネルを取得
func GetChannelsByWorkspaceID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// ワークスペースidを取得
		workspaceID := r.URL.Query().Get("workspace")

		// チャンネルを取得
		channels, err := usecase.GetChannelsByWorkspaceID(workspaceID)
		if err != nil {
			log.Printf("fail: usecase.GetChannelsByWorkspaceID, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// JSONに変換
		jsonBytes, err := json.Marshal(channels)
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
