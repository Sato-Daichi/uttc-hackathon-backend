package dao

import (
	"app/model"
	"log"
)

// userを登録する
func UserSignUp(user model.UserResForPost) (err error) {
	stmt, err := db.Prepare("INSERT INTO users (id, username, password, email) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Username, user.Password, user.Email)
	if err != nil {
		log.Printf("fail: stmt.Exec, %v\n", err)
		return err
	}
	return nil
}

// 全ユーザーを取得する
func GetAllUsers() ([]model.User, error) {
	stmt, err := db.Prepare("SELECT * FROM users")
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	log.Printf("rows: %v\n", rows)
	if err != nil {
		log.Printf("fail: stmt.Query, %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// emailとpasswordを受け取り、idとusernameを返す
func UserLogin(email string, password string) (string, string, error) {
	stmt, err := db.Prepare("SELECT id, username FROM users WHERE email = ? AND password = ?")
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return "", "", err
	}
	defer stmt.Close()

	var username string
	var userId string
	err = stmt.QueryRow(email, password).Scan(&userId, &username)
	if err != nil {
		log.Printf("fail: stmt.QueryRow.Scan, %v\n", err)
		return "", "", err
	}

	return userId, username, nil
}
