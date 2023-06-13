package controller

import (
	"app/model"
	"app/usecase"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/oklog/ulid"
)

// userを登録する
func UserSignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// jsonを解析してusername、password、emailを取得する
		var user model.UserResForPost
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("fail: json.NewDecoder.Decode, %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if user.Username == "" {
			log.Printf("fail: username is empty\n")
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if user.Password == "" {
			log.Printf("fail: password is empty\n")
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if user.Email == "" {
			log.Printf("fail: email is empty\n")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// uuidでidを生成する
		t := time.Now()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		user.Id = ulid.MustNew(ulid.Timestamp(t), entropy).String()

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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
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
