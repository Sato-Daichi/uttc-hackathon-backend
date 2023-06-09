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
		} else if messageResForPost.Text == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "text is empty")
			return
		} else if messageResForPost.ChannelId == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "channelId is empty")
			return
		} else if messageResForPost.UserId == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "userId is empty")
			return
		}

		// messageを作成
		// MessageResForPostからMessageに値を受け渡す
		var message model.PostMessage

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
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "success")
	case http.MethodOptions:
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
