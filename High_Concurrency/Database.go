package High_Concurrency

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "root"
	password = "123456"
	dbname   = "my_database"
)

var db *sql.DB
var err error

// 连接数据库
func Connect_database() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("无法连接数据库%v", err)
	}
	// 设置连接池
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(0)

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("<UNK>%v", err)
	}
	fmt.Println("数据库连接成功")
	return nil
}

// 插入数据
func InsertData(table string, data map[string]interface{}) (int64, error) {
	if db == nil {
		return 0, errors.New("数据库未连接，请先调用ConnectDatabase")
	}

	// 构建插入语句
	fields := ""
	placeholders := ""
	values := make([]interface{}, 0, len(data))
	i := 1

	for field, value := range data {
		if fields != "" {
			fields += ", "
			placeholders += ", "
		}
		fields += field
		placeholders += fmt.Sprintf("$%d", i)
		values = append(values, value)
		i++
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", table, fields, placeholders)

	var lastInsertID int64
	// 执行插入并返回自增ID
	err := db.QueryRow(query, values...).Scan(&lastInsertID)
	if err != nil {
		return 0, fmt.Errorf("插入数据失败: %v", err)
	}

	return lastInsertID, nil
}

// 删除数据
func DeleteData(table string, condition string, args ...interface{}) (int64, error) {
	if db == nil {
		return 0, errors.New("数据库未连接，请先调用ConnectDatabase")
	}

	if condition == "" {
		return 0, errors.New("删除条件不能为空，防止误删除所有数据")
	}

	// 构建删除语句
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", table, condition)

	// 执行删除
	result, err := db.Exec(query, args...)
	if err != nil {
		return 0, fmt.Errorf("删除数据失败: %v", err)
	}

	// 返回受影响的行数
	return result.RowsAffected()
}

// CloseDatabase 关闭数据库连接
func CloseDatabase() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
