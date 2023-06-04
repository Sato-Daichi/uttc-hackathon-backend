package usecase

import (
	"app/dao"
	"app/model"
)

// ユーザーidからワークスペースを取得
func GetWorkspacesByUserID(userID string) ([]model.Workspace, error) {
	workspaces, err := dao.GetWorkspacesByUserID(userID)
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}
