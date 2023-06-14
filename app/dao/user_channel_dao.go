package dao

import (
	"app/model"
	"log"
)

// ユーザーidとチャンネルidを紐づける
func JoinChannels(userChannels []model.UserChannel) error {
	// トランザクションを考慮
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	// ユーザーidとチャンネルidを紐づける
	for _, userChannel := range userChannels {
		_, err := tx.Exec("INSERT INTO users_channels (id, user_id, channel_id) VALUES (?, ?, ?)", userChannel.Id, userChannel.UserId, userChannel.ChannelId)
		if err != nil {
			log.Printf("fail: tx.Exec, %v\n", err)
			tx.Rollback()
			return err
		}
	}

	// コミット
	err = tx.Commit()
	if err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		return err
	}

	return nil
}
