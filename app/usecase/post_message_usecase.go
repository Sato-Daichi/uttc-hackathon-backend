package usecase

import (
	"app/dao"
	"app/model"
)

// messageを投稿
func PostMessage(message model.PostMessage) error {
	// messageを投稿
	err := dao.PostMessage(message)
	if err != nil {
		return err
	}

	return nil
}
