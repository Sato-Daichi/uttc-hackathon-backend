package usecase

import (
	"app/dao"
	"app/model"
	"log"
)

// userを登録する
func UserSignUp(user model.UserResForPost) error {
	// userを登録する
	err := dao.UserSignUp(user)
	if err != nil {
		log.Printf("fail: userDao.UserSignUp, %v\n", err)
		return err
	}

	return nil
}
