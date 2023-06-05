package usecase

import (
	"app/dao"
	"app/model"
)

// workspace_idから全チャンネルを取得
func GetChannelsByWorkspaceId(workspaceId string) ([]model.Channel, error) {
	channels, err := dao.GetChannelsByWorkspaceId(workspaceId)
	if err != nil {
		return nil, err
	}

	return channels, nil
}
