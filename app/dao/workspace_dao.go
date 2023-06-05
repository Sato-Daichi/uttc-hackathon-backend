package dao

import (
	"app/model"
	"log"
)

// user_idからworkspacesを取得
func GetWorkspacesByUserId(userId string) ([]model.Workspace, error) {
	stmt, err := db.Prepare("SELECT workspaces.id, workspaces.name, workspaces.created_at, workspaces.updated_at FROM workspaces LEFT JOIN users_workspaces ON workspaces.id = users_workspaces.workspace_id LEFT JOIN users ON users_workspaces.user_id = users.id WHERE users.id = ?")
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return nil, err
	}
	defer stmt.Close()

	row, err := stmt.Query(userId)
	if err != nil {
		log.Printf("fail: stmt.Query, %v\n", err)
		return nil, err
	}
	defer row.Close()

	var workspaces []model.Workspace
	for row.Next() {
		var workspace model.Workspace
		err = row.Scan(&workspace.Id, &workspace.Name, &workspace.CreatedAt, &workspace.UpdatedAt)
		if err != nil {
			log.Printf("fail: row.Scan, %v\n", err)
			return nil, err
		}
		workspaces = append(workspaces, workspace)
	}

	return workspaces, nil
}
