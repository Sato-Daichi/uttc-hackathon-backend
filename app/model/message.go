// Message構造体を定義する
// ただし、DBのmessagesテーブルの定義は以下の通り
// CREATE TABLE `messages` (
//   `id` varchar(255) NOT NULL,
//   `text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
//   `channel_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
//   `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
//   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//   `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//   PRIMARY KEY (`id`),
//   KEY `channel_id` (`channel_id`),
//   KEY `user_id` (`user_id`),
//   CONSTRAINT `messages_ibfk_1` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`id`),
//   CONSTRAINT `messages_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

package model

import (
	"time"
)

type Message struct {
	Id        string    `json:"id"`
	Text      string    `json:"text"`
	ChannelId string    `json:"channelId"`
	UserId    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PostMessage struct {
	Id 	  string `json:"id"`
	Text      string `json:"text"`
	ChannelId string `json:"channelId"`
	Username  string `json:"username"`
}
