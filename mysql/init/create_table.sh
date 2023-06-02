#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table app_user
(
    user_id   char(26)    not null
        primary key,
    user_name varchar(20) not null,
    password  varchar(50) not null
);"

$CMD_MYSQL -e "create table channel
(
    channel_id   char(26)    not null
        primary key,
    channel_name varchar(30) not null
);"

$CMD_MYSQL -e "create table channel_member
(
    id         char(26)    not null
        primary key,
    channel_id varchar(20) not null,
    user_id    char(26)    not null
);"

$CMD_MYSQL -e "create table message
(
    message_id      char(26)     not null
        primary key,
    user_id         char(26)     not null,
    message_content varchar(150) not null
);"

$CMD_MYSQL -e "create table reply
(
    reply_id      char(26)     not null
        primary key,
    user_id       char(26)     not null,
    reply_content varchar(150) not null,
    message_id    char(26)     not null
);"
