package usecase

import (
	"app/dao"
	"app/model"
)

// チャンネルを新規作成
func CreateChannel(channel model.Channel) (model.Channel, error) {
	channel, err := dao.CreateChannel(channel)
	if err != nil {
		return channel, err
	}

	return channel, nil
}
