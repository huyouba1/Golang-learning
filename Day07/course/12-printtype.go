package main

import (
	"fmt"
	"reflect"
	"time"
)

// 获取变量数据类型
func GetType(typ reflect.Type) string {
	//reflect.Type()
	switch typ.Kind() {
	case reflect.Int, reflect.String, reflect.Float64, reflect.Bool:
		return typ.Name()
	case reflect.Array:
		// [3]int 数组的数据类型 typ.Elem() 基本数据类型直接显示名称
		//
		return fmt.Sprintf("[%d]%s\n", typ.Len(), GetType(typ.Elem()))
	case reflect.Slice:
		return fmt.Sprintf("[]%s", GetType(typ.Elem()))
	case reflect.Map:
		return fmt.Sprintf("map[%s][%s]", GetType(typ.Key()), GetType(typ.Elem()))
	case reflect.Struct:
		return
	default:
		return "unknow"
	}
}

func main() {
	var es = []interface{}{1.2, 1, "test", false, [2]int{1, 2}, []int{1, 2, 3, 4}, map[int]string{1: "kk"}, time.Now()}

	// 打印每种变量的数据类型
	for _, e := range es {
		fmt.Println(GetType(reflect.TypeOf(e)))
	}

}
