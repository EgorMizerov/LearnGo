package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"math/rand"
	"time"
)

// db:"row_name" сопостовляет имя столбца с именем поля структуры
// db необзательно указывать явно, если название поли структуры в нижнем регистре будет соответствовать имени столбца в таблице
// sql.NullType указывет на то, что поле может хранить в себе указанынй Type или не хрнаить его
type post struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Text sql.NullString
}

func main() {
	// DB connection
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create Table
	// Exec просто выполняет SQL запрос
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Генерируем дату
	if err = addData(db); err != nil {
		log.Fatal(err.Error())
	}

	// Список строк
	rows, err := db.Queryx("SELECT * FROM posts")
	for rows.Next() {
		var x post
		// Скан строки в структуру
		err = rows.StructScan(&x)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println(x)
	}

	// Первая подходящая строка
	var x post
	err = db.QueryRowx("SELECT * FROM posts").StructScan(&x)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(x)

	// Первая подходящая строка
	err = db.Get(&x, "SELECT * FROM posts")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(x)

	// Список строк
	var xx []post
	err = db.Select(&xx, "SELECT * FROM posts LIMIT $1", 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(xx)

	// Транзакции
	tx, _ := db.Beginx()
	_, err = tx.Exec("INSERT INTO posts (name, text) VALUES ('Моя статья', 'Описание статьи')")
	if err != nil {
		// Rollback отменяет транзакцию
		tx.Rollback()
		log.Fatal(err.Error())
	}
	// Commit завершает транзакцию
	tx.Commit()

	// Шаблоны
	var p post
	stmt, _ := db.Preparex(`SELECT * FROM posts WHERE id=$1`)
	stmt.Get(&p, 12)
	fmt.Println(p)

	// Named Queries
	m := map[string]interface{}{"name": "Моя статья"}
	var xy post
	nstmt, err := db.PrepareNamed(`SELECT * FROM posts WHERE name=:name`)
	err = nstmt.Get(&xy, m)
	fmt.Println(xy)

	// Оператор IN
	var id = []int{4, 6, 7}
	query, args, err := sqlx.In("SELECT * FROM posts WHERE id IN ($1);", id)
	if err != nil {
		log.Fatal(err.Error())
	}

	query = db.Rebind(query)
	rows, err = db.Queryx(query, args...)

	for rows.Next() {
		var x post
		rows.StructScan(&x)
		fmt.Println(x)
	}

	// Drop Table
	db.Exec(dropTable)
}

var createTable = `
CREATE TABLE posts
(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    text TEXT
)`

var dropTable = "DROP TABLE posts"

func connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", "postgres://postgres:postgres@localhost:5432/sqlx?sslmode=disable")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func addData(db *sqlx.DB) error {
	for i := 1; i < 100; i++ {
		if i == 99 {
			query := fmt.Sprintf("INSERT INTO posts (name) VALUES ($1)")
			_, err := db.Exec(query, RandStringRunes(10))
			if err != nil {
				return err
			}
			continue
		}
		query := fmt.Sprintf("INSERT INTO posts (name, text) VALUES ($1, $2)")
		_, err := db.Exec(query, RandStringRunes(10), RandStringRunes(40))
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Навигация
/*
31 - Создание таблиц
43 - Список строк
56 - Первая строка
63 - Десереализация первой строк в структуру
71 - Десереализация списка строк в список структур
78 - Транзакции
88 - Шаблоны
95 - Именнованый поиск

*/
