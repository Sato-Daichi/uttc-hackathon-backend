package dao

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid/v2"
)

// テーブルと条件を指定して、条件を満たすレコードを1件取得する
// レコードを格納する構造体は引数で受け取る

// テーブルと条件を指定して、条件を満たすレコードを全件取得する
// レコードを格納する配列のポインタを引数で受け取る
func GetRecords(table string, record interface{}, where string, args ...interface{}) ([]interface{}, error) {
	stmt, err := db.Prepare(fmt.Sprintf("SELECT * FROM %s WHERE %s", table, where))
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		log.Printf("fail: stmt.Query, %v\n", err)
		return nil, err
	}

	var records []interface{}
	for rows.Next() {
		err = rows.Scan(record)
		if err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

// テーブルと構造体を引数として受け取り、レコードを1件挿入する
// idはULIdで生成する
// crated_atとupdated_atはmySQLの設定で自動で設定されるため、コーディングする必要はない
func InsertRecord(table string, record interface{}) error {
	stmt, err := db.Prepare(fmt.Sprintf("INSERT INTO %s VALUES ?", table))
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return err
	}
	defer stmt.Close()

	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy).Time()

	// レコードのIdを指定してレコードを挿入
	err = stmt.QueryRow(id, record).Scan()
	if err != nil {
		log.Printf("fail: stmt.QueryRow.Scan, %v\n", err)
		return err
	}

	return nil
}

// テーブルとレコードを指定して、レコードを1件更新する
// レコードを格納する構造体は引数で受け取る
func UpdateRecord(table string, record interface{}) error {
	stmt, err := db.Prepare(fmt.Sprintf("UPDATE %s SET ? WHERE ?", table))
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return err
	}
	defer stmt.Close()

	// レコードのIdを設定
	err = stmt.QueryRow(record).Scan()
	if err != nil {
		log.Printf("fail: stmt.QueryRow.Scan, %v\n", err)
		return err
	}

	return nil
}

// テーブルと条件を指定して、条件を満たすレコードを1件削除する
func DeleteRecord(table string, where string, args ...interface{}) error {
	stmt, err := db.Prepare(fmt.Sprintf("DELETE FROM %s WHERE %s", table, where))
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return err
	}
	defer stmt.Close()

	// レコードのIdを設定
	_, err = stmt.Exec(args...)
	if err != nil {
		log.Printf("fail: stmt.Exec, %v\n", err)
		return err
	}

	return nil
}
