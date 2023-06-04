package usecase

import (
	"app/dao"
	"app/model"
)

// チャンネルidから全メッセージを取得
func GetMessagesByChannelID(channelID string) ([]model.MessageUser, error) {
	messages, err := dao.GetMessagesByChannelID(channelID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
