package usecase

import (
	"app/dao"
	"app/model"
	"log"
)

// ユーザーidとチャンネルidを紐づける
func JoinChannels(userChannels []model.UserChannel) error {
	err := dao.JoinChannels(userChannels)
	if err != nil {
		log.Printf("fail: dao.JoinChannels, %v\n", err)
		return err
	}

	return nil
}
