package controller

import (
	"app/usecase"
	"encoding/json"
	"log"
	"net/http"
)

// ユーザーidからワークスペースを取得
func GetWorkspacesByUserId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// ユーザーidを取得
		userId := r.URL.Query().Get("user")

		// ワークスペースを取得
		workspaces, err := usecase.GetWorkspacesByUserId(userId)
		if err != nil {
			log.Printf("fail: usecase.GetWorkspacesByUserId, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// JSONに変換
		jsonBytes, err := json.Marshal(workspaces)
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
