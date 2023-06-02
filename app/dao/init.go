package dao

import (
	"database/sql"
	"log"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/joho/godotenv"
)
var db *sql.DB
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("fail: godotenv.Load, %v\n", err)
	}

	mysqlUser := os.Getenv("MYSQL_USER")
    mysqlPwd := os.Getenv("MYSQL_PWD")
    mysqlHost := os.Getenv("MYSQL_HOST")
    mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	// DB接続
	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
    db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// 実際に接続できるか確認
	if err := db.Ping(); err != nil {
		log.Fatalf("fail: db.Ping, %v\n", err)
	}

	closeDBWithSysCall()
}

// Ctrl+CでHTTPサーバー停止時にDBをクローズする
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
