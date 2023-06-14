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

// userIdとchannelIdsを紐づける
func JoinChannels(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// リクエストボディをパース
		type RequestBody struct {
			UserId     string   `json:"userId"`
			ChannelIds []string `json:"channelIds"`
		}

		var requestBody RequestBody
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("fail: json.NewDecoder().Decode, %v\n", err)
			return
		} else if requestBody.UserId == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "userId is empty")
			return
		} else if requestBody.ChannelIds == nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "channelIds is empty")
			return
		}

		// userChannelの配列を作成する
		var userChannels []model.UserChannel
		for _, channelId := range requestBody.ChannelIds {
			// idはuuidで生成
			t := time.Now()
			entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
			Id := ulid.MustNew(ulid.Timestamp(t), entropy).String()

			userChannel := model.UserChannel{
				Id:        Id,
				UserId:    requestBody.UserId,
				ChannelId: channelId,
			}

			userChannels = append(userChannels, userChannel)
		}

		// userIdとchannelIdsを紐づける
		err = usecase.JoinChannels(userChannels)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("fail: usecase.JoinChannels, %v\n", err)
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
