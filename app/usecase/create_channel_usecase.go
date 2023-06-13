package usecase

import (
	"app/dao"
	"app/model"
)

// チャンネルを新規作成
func CreateChannel(channel model.Channel) error {
	// チャンネルを作成
	if err := dao.CreateChannel(channel); err != nil {
		return err
	}

	return nil
}
