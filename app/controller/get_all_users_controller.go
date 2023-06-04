package controller

import (
	"app/usecase"
	"encoding/json"
	"log"
	"net/http"
)

// 全ユーザーを取得
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// ユーザーを取得
	users, err := usecase.GetAllUsers()
	if err != nil {
		log.Printf("fail: usecase.GetAllUsers, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// JSONに変換
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// レスポンスに書き込み
	w.Write(jsonBytes)
}