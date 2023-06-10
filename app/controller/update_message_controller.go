package controller

import (
	"app/model"
	"app/usecase"
	"encoding/json"
	"fmt"
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
