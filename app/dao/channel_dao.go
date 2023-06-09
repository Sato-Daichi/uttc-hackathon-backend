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

// workspace_idとuser_idから全チャンネルを取得
func GetChannelsByWorkspaceIdAndUserId(workspaceId string, userId string) ([]model.Channel, error) {
	stmt, err := db.Prepare("SELECT channels.id, channels.name, channels.description, channels.create_user_id, channels.created_at, channels.updated_at, channels.workspace_id FROM users LEFT JOIN users_channels ON users.id = users_channels.user_id LEFT JOIN channels ON users_channels.channel_id = channels.id WHERE channels.workspace_id = ? AND users.id = ?")
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(workspaceId, userId)
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

// チャンネルを新規作成
func CreateChannel(channel model.Channel) (model.Channel, error) {
	stmt, err := db.Prepare("INSERT INTO channels (id, name, description, create_user_id, workspace_id) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return channel, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(channel.Id, channel.Name, channel.Description, channel.CreateUserId, channel.WorkspaceId)
	if err != nil {
		log.Printf("fail: stmt.Exec, %v\n", err)
		return channel, err
	}

	// idをもとにチャンネルを取得
	stmt, err = db.Prepare("SELECT created_at, updated_at FROM channels WHERE id = ?")
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return channel, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(channel.Id).Scan(&channel.CreatedAt, &channel.UpdatedAt)
	if err != nil {
		log.Printf("fail: stmt.QueryRow.Scan, %v\n", err)
		return channel, err
	}

	return channel, nil
}
