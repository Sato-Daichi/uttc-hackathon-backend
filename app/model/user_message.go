package model

import (
	"time"
)

// messageとuserをjoinした構造体を定義する
type MessageUser struct {
	Id            string    `json:"id"`
	Text          string    `json:"text"`
	ChannelId     string    `json:"channelId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UserId        string    `json:"userId"`
	UserUsername  string    `json:"userUsername"`
	UserPassword  string    `json:"userPassword"`
	UserEmail     string    `json:"userEmail"`
	UserCreatedAt time.Time `json:"userCreatedAt"`
	UserUpdatedAt time.Time `json:"userUpdatedAt"`
}
