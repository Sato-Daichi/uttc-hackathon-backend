package usecase

import (
	"app/dao"
	"app/model"
)

// チャンネルidから全メッセージを取得
func GetMessagesByChannelId(channelId string) ([]model.MessageUser, error) {
	messages, err := dao.GetMessagesByChannelId(channelId)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
