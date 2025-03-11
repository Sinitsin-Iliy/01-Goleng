package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB представляет объект соединения с базой данных
var DB *sql.DB

// ConnectToDatabase устанавливает соединение с базой данных
func ConnectToDatabase() {
	dsn := "zakonm_user:zakonm_password@tcp(db:3306)/zakonm_db?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Ошибка пинга базы данных: %v", err)
	}
	fmt.Println("Соединение с базой данных установлено!")
}
