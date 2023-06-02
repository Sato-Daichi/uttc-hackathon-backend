// UserWorkspace構造体を定義する
// ただし、DBのusers_workspacesテーブルの定義は以下の通り
// CREATE TABLE `users_workspaces` (
//   `id` varchar(255) NOT NULL,
//   `workspace_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
//   `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
//   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//   `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//   PRIMARY KEY (`id`),
//   KEY `workspace_id` (`workspace_id`),
//   KEY `user_id` (`user_id`),
//   CONSTRAINT `users_workspaces_ibfk_1` FOREIGN KEY (`workspace_id`) REFERENCES `workspaces` (`id`),
//   CONSTRAINT `users_workspaces_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

package model

import (
	"time"
)

type UserWorkspace struct {
	Id          string    `json:"id"`
	WorkspaceId string    `json:"workspace_id"`
	UserId      string    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}