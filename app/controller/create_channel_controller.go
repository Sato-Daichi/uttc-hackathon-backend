package controller

import (
	"app/model"
	"app/usecase"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/oklog/ulid"
)

// CreateChannel はチャンネルを作成する
func CreateChannel(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// リクエストボディをパース
		var channel model.Channel
		if err := json.NewDecoder(r.Body).Decode(&channel); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// idはuuidで生成
		t := time.Now()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		channel.Id = ulid.MustNew(ulid.Timestamp(t), entropy).String()

		// チャンネルを作成
		if err := usecase.CreateChannel(channel); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// レスポンスを返す
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
