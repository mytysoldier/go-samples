package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=user password=password dbname=sample sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB接続")

	// コマンドラインからの入力を受け付ける
	var command string
	fmt.Print("コマンドを入力 (insert/update/delete/select): ")
	fmt.Scan(&command)

	switch command {
	case "insert":
		dbInsert(db)
	case "update":
		dbUpdate(db)
	case "select":
		dbSelect(db)
	case "delete":
		dbDelete(db)
	default:
		fmt.Println("無効なコマンド")
	}
}

func dbInsert(db *sql.DB) {
	var name string
	var age int

	fmt.Print("名前を入力: ")
	fmt.Scan(&name)

	fmt.Print("年齢を入力: ")
	fmt.Scan(&age)

	// INSERT文を実行
	result, err := db.Exec("INSERT INTO \"user\" (name, age) VALUES ($1, $2)", name, age)
	if err != nil {
		log.Fatal(err)
	}

	// 追加成功時のメッセージと追加されたレコード内容を表示
	fmt.Println("追加しました。")
	lastInsertID, _ := result.LastInsertId()
	fmt.Printf("追加されたレコードのID: %d, 名前: %s, 年齢: %d\n", lastInsertID, name, age)

}

func dbUpdate(db *sql.DB) {
	var id int
	var name string
	var age int

	fmt.Print("更新対象のレコードのIDを入力: ")
	fmt.Scan(&id)

	fmt.Print("名前を入力: ")
	fmt.Scan(&name)

	fmt.Print("年齢を入力: ")
	fmt.Scan(&age)

	// UPDATE文を実行
	result, err := db.Exec("UPDATE \"user\" SET name = $1, age = $2 WHERE id = $3", name, age, id)
	if err != nil {
		log.Fatal(err)
	}

	// 更新成功時のメッセージと更新されたレコード内容を表示
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected > 0 {
		fmt.Println("更新しました。")
		fmt.Printf("更新されたレコードのID: %d, 新しい名前: %s, 新しい年齢: %d\n", id, name, age)
	} else {
		fmt.Println("指定されたIDのレコードが見つかりませんでした。")
	}
}

func dbSelect(db *sql.DB) {
	var id int

	fmt.Print("検索対象のレコードのIDを入力: ")
	fmt.Scan(&id)

	// SELECT文を実行
	row := db.QueryRow("SELECT id, name, age FROM \"user\" WHERE id = $1", id)

	var selectedID int
	var selectedName string
	var selectedAge int

	// 取得したレコードの値をスキャン
	err := row.Scan(&selectedID, &selectedName, &selectedAge)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("指定されたIDのレコードが見つかりませんでした。")
		} else {
			log.Fatal(err)
		}
		return
	}

	// 取得したレコードの内容を表示
	fmt.Printf("ID: %d, 名前: %s, 年齢: %d\n", selectedID, selectedName, selectedAge)
}

func dbDelete(db *sql.DB) {
	var id int

	fmt.Print("削除対象のレコードのIDを入力: ")
	fmt.Scan(&id)

	// DELETE文を実行
	result, err := db.Exec("DELETE FROM \"user\" WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	// 削除成功時のメッセージを表示
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected > 0 {
		fmt.Println("削除しました。")
	} else {
		fmt.Println("指定されたIDのレコードが見つかりませんでした。")
	}
}
