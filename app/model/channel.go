// Channel構造体を定義する
// ただし、DBのchannelsテーブルの定義は以下の通り
// CREATE TABLE `channels` (
//   `id` varchar(255) NOT NULL,
//   `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
//   `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
//   `create_user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
//   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//   `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//   `workspace_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
//   PRIMARY KEY (`id`),
//   KEY `create_user_id` (`create_user_id`),
//   KEY `workspace_id` (`workspace_id`),
//   CONSTRAINT `channels_ibfk_1` FOREIGN KEY (`create_user_id`) REFERENCES `users` (`id`),
//   CONSTRAINT `channels_ibfk_2` FOREIGN KEY (`workspace_id`) REFERENCES `workspaces` (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

package model

import (
	"time"
)

type Channel struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreateUserId string    `json:"createUserId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	WorkspaceId  string    `json:"workspaceId"`
}
