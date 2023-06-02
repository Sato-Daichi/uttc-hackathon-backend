#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"

$CMD_MYSQL -e "CREATE TABLE channels (
  id varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  description text NOT NULL,
  create_user_id varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  workspace_id varchar(255) NOT NULL
);"

$CMD_MYSQL -e "CREATE TABLE messages (
  id varchar(255) NOT NULL,
  text text NOT NULL,
  channel_id varchar(255) NOT NULL,
  user_id varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);"

$CMD_MYSQL -e "CREATE TABLE replies (
  id varchar(255) NOT NULL,
  text text NOT NULL,
  channel_id varchar(255) NOT NULL,
  user_id varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  parent_id varchar(255) NOT NULL
);"

$CMD_MYSQL -e "CREATE TABLE users_channels (
  id varchar(255) NOT NULL,
  channel_id varchar(255) NOT NULL,
  user_id varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);"

$CMD_MYSQL -e "CREATE TABLE users_workspaces (
  id varchar(255) NOT NULL,
  workspace_id varchar(255) NOT NULL,
  user_id varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);"

$CMD_MYSQL -e "CREATE TABLE users (
  id varchar(26) NOT NULL,
  username varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);"

$CMD_MYSQL -e "CREATE TABLE workspaces (
  id varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);"

# ユーザーを作成
$CMD_MYSQL -e "INSERT INTO users (id, username, password, email) VALUES ('00000000000000000000000001', 'sato', 'sato', 'sato@gmail.com');"
$CMD_MYSQL -e "INSERT INTO users (id, username, password, email) VALUES ('00000000000000000000000002', 'suzuki', 'suzuki', 'suzuki@gmail.com');"

# ワークスペースを作成
$CMD_MYSQL -e "INSERT INTO workspaces (id, name) VALUES ('00000000000000000000000001', 'UTokyo Tech Club');"

# チャンネルを作成
$CMD_MYSQL -e "INSERT INTO channels (id, name, description, create_user_id, workspace_id) VALUES ('00000000000000000000000001', 'general', '一般的な内容', '00000000000000000000000001', '00000000000000000000000001');"

# メッセージを作成
$CMD_MYSQL -e "INSERT INTO messages (id, text, channel_id, user_id) VALUES ('00000000000000000000000001', 'こんにちは！', '00000000000000000000000001', '00000000000000000000000001');"
$CMD_MYSQL -e "INSERT INTO messages (id, text, channel_id, user_id) VALUES ('00000000000000000000000002', 'こんにちは！', '00000000000000000000000001', '00000000000000000000000002');"

# チャンネルとユーザーを紐付け
$CMD_MYSQL -e "INSERT INTO users_channels (id, channel_id, user_id) VALUES ('00000000000000000000000001', '00000000000000000000000001', '00000000000000000000000001');"
$CMD_MYSQL -e "INSERT INTO users_channels (id, channel_id, user_id) VALUES ('00000000000000000000000002', '00000000000000000000000001', '00000000000000000000000002');"

# ワークスペースとユーザーを紐付け
$CMD_MYSQL -e "INSERT INTO users_workspaces (id, workspace_id, user_id) VALUES ('00000000000000000000000001', '00000000000000000000000001', '00000000000000000000000001');"
$CMD_MYSQL -e "INSERT INTO users_workspaces (id, workspace_id, user_id) VALUES ('00000000000000000000000002', '00000000000000000000000001', '00000000000000000000000002');"