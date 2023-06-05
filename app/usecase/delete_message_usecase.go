package usecase

import (
	"app/dao"
	"log"
)

// message_idからメッセージを削除する
func DeleteMessage(messageId string) error {
	err := dao.DeleteMessage(messageId)
	if err != nil {
		log.Printf("fail: dao.DeleteMessage, %v\n", err)
		return err
	}
	return nil
}
