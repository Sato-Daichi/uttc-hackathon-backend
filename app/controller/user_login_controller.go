package controller

import (
	"app/model"
	"app/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

// emailとpasswordを受け取り、usernameを返す
func UserLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		email := r.URL.Query().Get("email")
		password := r.URL.Query().Get("password")

		if email == "" {
			fmt.Println("fail: r.URL.Query().Get, email")
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if password == "" {
			fmt.Println("fail: r.URL.Query().Get, password")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// userを登録する
		userId, username, err := usecase.UserLogin(email, password)
		if err != nil {
			fmt.Println("fail: usecase.UserLogin,", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// レスポンスボディを作成
		res := model.UserResForPost{
			Id:       userId,
			Username: username,
		}

		// レスポンスボディを作成
		resBody, err := json.Marshal(res)
		if err != nil {
			fmt.Println("fail: json.Marshal(res),", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// レスポンスヘッダを作成
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
		// レスポンスボディを書き込む
		w.Write(resBody)
	case http.MethodOptions:
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
