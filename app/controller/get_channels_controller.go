package controller

import (
	"app/usecase"
	"encoding/json"
	"log"
	"net/http"
)

// ワークスペースidからチャンネルを取得
func GetChannelsByWorkspaceId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// ワークスペースidを取得
		workspaceId := r.URL.Query().Get("workspace")
		if workspaceId == "" {
			log.Printf("fail: r.URL.Query().Get, %v\n", workspaceId)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// チャンネルを取得
		channels, err := usecase.GetChannelsByWorkspaceId(workspaceId)
		if err != nil {
			log.Printf("fail: usecase.GetChannelsByWorkspaceId, %v\n", err)
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
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	case http.MethodOptions:
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
