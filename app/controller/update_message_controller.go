package controller

import (
	"app/model"
	"app/usecase"
	"encoding/json"
	"log"
	"net/http"
)

// messageをupdateする
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPatch:
		// リクエストボディを取得
		var messageResForPatch model.MessageResForPatch
		err := json.NewDecoder(r.Body).Decode(&messageResForPatch)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("fail: json.NewDecoder().Decode, %v\n", err)
			return
		}

		// messageをupdate
		err = usecase.UpdateMessage(messageResForPatch)
		if err != nil {
			log.Printf("fail: usecase.UpdateMessage, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
