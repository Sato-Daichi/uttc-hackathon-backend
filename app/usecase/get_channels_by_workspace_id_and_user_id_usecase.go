package usecase

import (
	"app/dao"
	"app/model"
)

// workspace_idから全チャンネルを取得
func GetChannelsByWorkspaceIdAndUserId(workspaceId string, userId string) ([]model.Channel, error) {
	channels, err := dao.GetChannelsByWorkspaceIdAndUserId(workspaceId, userId)
	if err != nil {
		return nil, err
	}

	return channels, nil
}
