package dao

import (
	"app/model"
	"log"
)

// チャンネルidから全メッセージを取得
func GetMessagesByChannelID(channelID string) ([]model.MessageUser, error) {
	stmt, err := db.Prepare("SELECT messages.id, messages.text, messages.channel_id, messages.created_at, messages.updated_at, users.id, users.username, users.password, users.email, users.created_at, users.updated_at FROM messages INNER JOIN users ON messages.user_id = users.id WHERE messages.channel_id = ?")
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(channelID)
	if err != nil {
		log.Printf("fail: stmt.Query, %v\n", err)
		return nil, err
	}

	var messageUsers []model.MessageUser
	for rows.Next() {
		var messageUser model.MessageUser
		err = rows.Scan(&messageUser.Id, &messageUser.Text, &messageUser.ChannelId, &messageUser.CreatedAt, &messageUser.UpdatedAt, &messageUser.UserId, &messageUser.UserUsername, &messageUser.UserPassword, &messageUser.UserEmail, &messageUser.UserCreatedAt, &messageUser.UserUpdatedAt)
		if err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return nil, err
		}
		messageUsers = append(messageUsers, messageUser)
	}

	return messageUsers, nil
}
