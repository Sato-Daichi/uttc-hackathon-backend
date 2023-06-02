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

$CMD_MYSQL -e "INSERT INTO users (id, username, password, email) VALUES ('00000000000000000000000001', 'sato', 'sato', 'sato@gmail.com');"
$CMD_MYSQL -e "INSERT INTO users (id, username, password, email) VALUES ('00000000000000000000000002', 'suzuki', 'suzuki', 'suzuki@gmail.com');"
