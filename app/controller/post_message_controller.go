package controller

import (
	"app/model"
	"app/usecase"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/oklog/ulid"
)

type MessageResForPost struct {
	Text      string `json:"text"`
	ChannelId string `json:"channelId"`
	UserId    string `json:"userId"`
}

// messageを投稿
func PostMessage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// リクエストボディを取得
		var messageResForPost MessageResForPost
		err := json.NewDecoder(r.Body).Decode(&messageResForPost)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("fail: json.NewDecoder().Decode, %v\n", err)
			return
		}

		// messageを作成
		// MessageResForPostからMessageに値を受け渡す
		var message model.Message

		// uuidでIdを生成
		t := time.Now()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		message.Id = ulid.MustNew(ulid.Timestamp(t), entropy).String()
		message.Text = messageResForPost.Text
		message.ChannelId = messageResForPost.ChannelId
		message.UserId = messageResForPost.UserId

		// messageを投稿
		err = usecase.PostMessage(message)
		if err != nil {
			log.Printf("fail: usecase.PostMessage, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// レスポンスに書き込み
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "success")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
