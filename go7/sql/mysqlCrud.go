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
	res, err := db.Exec("insert into student (name,province,city,enrollment) values ('小明','深圳','深圳','2023-08-08')")
	if err != nil {
		log.Println("Insert sql failed: ", err)
		return
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

func update(db *sql.DB) {
	res, err := db.Exec("update student  set score=score+10 where city='深圳'")
	if err != nil {
		log.Println("update sql failed: ", err)
		return
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
	rows, err := db.Query("select id,score,name,city from student")
	if err != nil {
		log.Println("query sql failed: ", err)
		return
	}
	for rows.Next() {
		var id int
		var score float32
		var name, city string
		err := rows.Scan(&id, &score, &name, &city) // 通过scan把db里数据赋值给变量
		if err != nil {
			log.Println("result scan failed: ", err)
			return
		}
		log.Printf("id = %d, score = %.2f, name = %s, city = %s", id, score, name, city)
	}
}
