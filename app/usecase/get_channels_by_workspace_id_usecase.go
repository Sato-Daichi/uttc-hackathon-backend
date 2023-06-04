package usecase

import (
	"app/dao"
	"app/model"
)

// workspace_idから全チャンネルを取得
func GetChannelsByWorkspaceID(workspaceID string) ([]model.Channel, error) {
	channels, err := dao.GetChannelsByWorkspaceID(workspaceID)
	if err != nil {
		return nil, err
	}

	return channels, nil
}
