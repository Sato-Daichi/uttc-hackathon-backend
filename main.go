package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"os/signal"
	"syscall"
	"math/rand"

	"github.com/oklog/ulid/v2"
	_ "github.com/go-sql-driver/mysql"
)

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResForHTTPPost struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResWithId struct {
	Id   int `json:"id"`
}

// ① GoプログラムからMySQLへ接続
var db *sql.DB

// 最初に呼ばれる関数
func init() {
    mysqlUser := os.Getenv("MYSQL_USER")
    mysqlPwd := os.Getenv("MYSQL_PWD")
    mysqlHost := os.Getenv("MYSQL_HOST")
    mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	// DB接続
	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
    _db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// 実際に接続できるか確認
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	case http.MethodGet:
		// nameパラメータを取得
		name := r.URL.Query().Get("name")
		if name == "" {
			log.Println("fail: name is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// nameで検索
		rows, err := db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		// 1行ずつ取得
		users := make([]UserResForHTTPGet, 0)
		for rows.Next() {
			var u UserResForHTTPGet
			if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)

				if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		// jsonに変換
		bytes, err := json.Marshal(users)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	case http.MethodPost:
		// POSTされたらDBにnameとageをinsertする
		user := UserResForHTTPPost{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("fail: json.NewDecoder().Decode, %v\n", err)
			return
		}

		// バリデーション
		// nameが空文字, nameが50字より長い, ageが20歳未満, ageが80歳より大きい場合は400を返す
		if user.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("fail: name is empty\n")
			return
		} else if len(user.Name) > 50 {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("fail: name is too long, %s\n", user.Name)
			return
		}

		if user.Age < 20 || user.Age > 80 {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("fail: age is invalid, %d\n", user.Age)
			return
		}

		// トランザクション開始
		tx, err := db.Begin()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		// id作成
		t := time.Now()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		id := ulid.MustNew(ulid.Timestamp(t), entropy).Time()

		// insert	
		_, err = tx.Exec("INSERT INTO user (id, name, age) VALUES(?, ?, ?)", id, user.Name, user.Age)
		if err != nil {
			tx.Rollback()
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("failed: insert data to database, %v\n", err)
			return
		}
		log.Printf("success: insert data to database, \n")

		tx.Commit()
		
		// idのみ取り出す
		res := UserResWithId{
			Id: int(id),
		}

		// jsonに変換
		bytes, err := json.Marshal(res)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// usersでリクエストされたらDBのuserテーブルの全レコードをJSON形式で返す
	case http.MethodGet:
		// 全件取得
		rows, err := db.Query("SELECT id, name, age FROM user")
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// 1行ずつ取得
		users := make([]UserResForHTTPGet, 0)
		for rows.Next() {
			var u UserResForHTTPGet
			if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)

				if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		// jsonに変換
		bytes, err := json.Marshal(users)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(bytes)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {
	// usersでリクエストされたらDBのuserテーブルの全レコードをJSON形式で返す
	http.HandleFunc("/users", usersHandler)

	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", handler)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}