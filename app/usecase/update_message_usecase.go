package usecase

import (
	"app/dao"
	"app/model"
)

// messageをupdateする
func UpdateMessage(messageResForPatch model.MessageResForPatch) error {
	err := dao.UpdateMessage(messageResForPatch)
	if err != nil {
		return err
	}
	return nil
}
