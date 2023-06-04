package usecase

import (
	"app/dao"
	"log"
	"app/model"
)

// 全ユーザーを取得
func GetAllUsers() ([]model.User, error) {
	// ユーザーを取得
	users, err := dao.GetAllUsers()
	if err != nil {
		log.Printf("fail: dao.GetAllUsers, %v\n", err)
		return nil, err
	}

	return users, nil
}