package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"app/model"
	"app/usecase"
)

// userを登録する
func UserSignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// jsonを解析してusername、password、emailを取得する
		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user);
		if err != nil {
			log.Printf("fail: json.NewDecoder.Decode, %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// username、password、emailが空文字の場合は400エラーを返す
		if user.Username == "" || user.Password == "" || user.Email == "" {
			log.Printf("fail: empty username or password or email\n")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// userを登録する
		err = usecase.UserSignUp(user)
		if err != nil {
			log.Printf("fail: userSignUpService.UserSignUp, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// userをjsonに変換する
		jsonBytes, err := json.Marshal(user)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// jsonをレスポンスする
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
