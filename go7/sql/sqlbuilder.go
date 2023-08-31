package main

import (
	"fmt"
	gsb "github.com/parkingwang/go-sqlbuilder"
)

func main() {
	sql := gsb.NewContext().Insert("student").
		Columns("name", "province", "city", "enrollment").
		Values("赵丽", "江苏", "南京", "2021-07-09").ToSQL()
	fmt.Println(sql)
}
