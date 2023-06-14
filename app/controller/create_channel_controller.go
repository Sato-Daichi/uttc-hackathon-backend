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
		} else if channel.Name == "" {
			fmt.Println("name is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if channel.CreateUserId == "" {
			fmt.Println("createUserId is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if channel.WorkspaceId == "" {
			fmt.Println("workspaceId is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// idはuuidで生成
		t := time.Now()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		channel.Id = ulid.MustNew(ulid.Timestamp(t), entropy).String()

		// チャンネルを作成
		channel, err := usecase.CreateChannel(channel)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// jsonに変換
		jsonBytes, err := json.Marshal(channel)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// レスポンスを返す
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
