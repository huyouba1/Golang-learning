package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	//insert(db)
	//update(db)
	query(db)
	db.Close()
}

func insert(db *sql.DB) {
	stmt, err := db.Prepare("insert into student (name,province,city,enrollment) values (?,?,?,?),(?,?,?,?)")
	if err != nil {
		log.Println("Insert sql failed: ", err)
		return
	}

	//time.ParseInLocation()

	res, err := stmt.Exec("小亮", "深圳", "深圳", "2023-07-09", "大圈", "北京", "海淀", "2023-07-05")
	if err != nil {
		log.Println("exec failed ", err)
	}
	lastId, err := res.LastInsertId() // ID自增，用过的id(即使对应的行已delete)
	if err != nil {
		log.Println(err)
	}
	log.Printf("after insert last ID %d\n", lastId)
	rows, err := res.RowsAffected() // 影响了几行
	if err != nil {
		log.Println(err)
	}
	log.Printf("insert affect %d rows", rows)
}

func query(db *sql.DB) {
	stmt, err := db.Prepare("select id,name,city,score from student where id >?")
	if err != nil {
		log.Println("query stmt pare failed ", err)
	}
	rows, err := stmt.Query(2)
	for rows.Next() {
		var id int
		var name, city string
		var score float64
		err := rows.Scan(&id, &name, &city, &score)
		if err != nil {
			log.Println("scan failed", err)
		}
		log.Printf("id = %d, score = %.2f, name = %s, city = %s", id, score, name, city)
	}
}
