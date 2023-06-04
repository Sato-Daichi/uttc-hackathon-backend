package dao

import (
	"app/model"
	"log"
)

// userを登録する
func UserSignUp(user model.User) (model.User, error) {
	stmt, err := db.Prepare("INSERT INTO user (id, username, password, email) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return model.User{}, err
	}
	defer stmt.Close()

	row, err := stmt.Query(user.Id, user.Username, user.Password, user.Email)
	if err != nil {
		log.Printf("fail: stmt.Query, %v\n", err)
		return model.User{}, err
	}
	defer row.Close()

	if !row.Next() {
		log.Printf("fail: row.Next, %v\n", err)
		return model.User{}, err
	}

	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	if err != nil {
		log.Printf("fail: row.Scan, %v\n", err)
		return model.User{}, err
	}

	return user, nil
}

// emailを指定してuserを取得する
func GetUserByEmail(email string) (model.User, error) {
	stmt, err := db.Prepare("SELECT * FROM user WHERE email = ?")
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return model.User{}, err
	}
	defer stmt.Close()

	row, err := stmt.Query(email)
	if err != nil {
		log.Printf("fail: stmt.Query, %v\n", err)
		return model.User{}, err
	}
	defer row.Close()

	if !row.Next() {
		log.Printf("fail: row.Next, %v\n", err)
		return model.User{}, err
	}

	var user model.User
	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	if err != nil {
		log.Printf("fail: row.Scan, %v\n", err)
		return model.User{}, err
	}

	return user, nil
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
