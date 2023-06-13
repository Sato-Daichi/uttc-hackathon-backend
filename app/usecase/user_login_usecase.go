package usecase

import (
	"app/dao"
	"log"
)

// emailとpasswordを受け取り、usernameを返す
func UserLogin(email string, password string) (string, string, error) {
	// userを登録する
	userId, username, err := dao.UserLogin(email, password)
	if err != nil {
		log.Printf("fail: userDao.UserLogin, %v\n", err)
		return "", "", err
	}

	return userId, username, nil
}
