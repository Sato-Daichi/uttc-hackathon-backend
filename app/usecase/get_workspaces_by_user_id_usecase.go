package usecase

import (
	"app/dao"
	"app/model"
)

// ユーザーidからワークスペースを取得
func GetWorkspacesByUserId(userId string) ([]model.Workspace, error) {
	workspaces, err := dao.GetWorkspacesByUserId(userId)
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}
