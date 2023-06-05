package dao

import (
	"app/model"
	"log"
)

// workspace_idから全チャンネルを取得
func GetChannelsByWorkspaceId(workspaceId string) ([]model.Channel, error) {
	stmt, err := db.Prepare("SELECT channels.id, channels.name, channels.description, channels.create_user_id, channels.created_at, channels.updated_at, channels.workspace_id FROM channels LEFT JOIN workspaces ON channels.workspace_id = workspaces.id WHERE workspaces.id = ?")
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(workspaceId)
	if err != nil {
		log.Printf("fail: stmt.Query, %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var channels []model.Channel
	for rows.Next() {
		var channel model.Channel
		err = rows.Scan(&channel.Id, &channel.Name, &channel.Description, &channel.CreateUserId, &channel.CreatedAt, &channel.UpdatedAt, &channel.WorkspaceId)
		if err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return nil, err
		}
		channels = append(channels, channel)
	}

	return channels, nil
}
