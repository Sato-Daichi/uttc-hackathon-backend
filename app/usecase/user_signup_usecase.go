package usecase

import (
	"app/dao"
	"app/model"
	"database/sql"
	"fmt"
	"log"
)

// userを登録する
// ただしemailが既に登録されているかチェックする
func UserSignUp(user model.User) error {
	// emailが既に登録されている場合は400エラーを返す
	_, err := dao.GetUserByEmail(user.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("fail: userDao.GetUserByEmail, %v\n", err)
			return err
		}
	} else {
		log.Printf("fail: email already exists\n")
		return fmt.Errorf("email already exists")
	}

	// userを登録する
	user, err = dao.UserSignUp(user)
	if err != nil {
		log.Printf("fail: userDao.UserSignUp, %v\n", err)
		return err
	}

	return nil
}