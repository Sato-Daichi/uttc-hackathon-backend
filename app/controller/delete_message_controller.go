package controller

import (
	"app/usecase"
	"log"
	"net/http"
)

// message_idからメッセージを削除する
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		message_id := r.URL.Query().Get("message")
		err := usecase.DeleteMessage(message_id)
		if err != nil {
			log.Printf("fail: usecase.DeleteMessage, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
