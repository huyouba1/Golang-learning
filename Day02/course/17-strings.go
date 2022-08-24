package main

import (
	"fmt"
	"strings"
)

func main() {
	// 比较
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))

	// 包含
	fmt.Println(strings.Contains("我是xxx", "xxx"))
	fmt.Println(strings.Contains("我是xxx", "xxx1"))

	fmt.Println(strings.ContainsAny("我是xxx", "xxx1")) // 包含其中任意字符为true
	fmt.Println(strings.ContainsAny("我是xxx", "123"))

	fmt.Println(strings.ContainsRune("我是xxx", '我')) // rune格式，匹配Unicode
	fmt.Println(strings.ContainsRune("我是xxx", 'a'))

	//统计字符出现的次数
	fmt.Println(strings.Count("我是xxx", "xxx"))
	fmt.Println(strings.Count("我是xxx", "x"))
	fmt.Println(strings.Count("我是xxx", "a"))

	// 不区分大小写比较
	fmt.Println(strings.EqualFold("abc", "ABC"))
	fmt.Println(strings.EqualFold("abc", "abc"))
	fmt.Println(strings.EqualFold("abc", "xyz"))

	// 空白符分割
	// 空格 tab 回车 换行 换页 ...
	fmt.Printf("%#v\n", strings.Fields("aafds b\tc\nd\re\ff"))

	// 开头 * 结尾
	fmt.Println(strings.HasPrefix("abc", "ab"))
	fmt.Println(strings.HasPrefix("abc", "bc"))

	fmt.Println(strings.HasSuffix("abc", "ab"))
	fmt.Println(strings.HasSuffix("abc", "bc"))

	// 字符串出现位置
	fmt.Println(strings.Index("abcdefdef", "def")) // 匹配第一次出现的位置
	fmt.Println(strings.Index("abcdefdef", "xxx"))
	fmt.Println(strings.LastIndex("abcdefdef", "def")) // 匹配最后一次出现的位置
	fmt.Println(strings.LastIndex("abcdefdef", "xxx"))

	// 连接 分割
	fmt.Println(strings.Join([]string{"ab", "cd", "ef"}, "-"))
	fmt.Printf("%#v\n", strings.Split("ab-cd-ef", "-"))
	fmt.Printf("%#v\n", strings.SplitN("ab-cd-ef", "-", 2))

	// 重复
	fmt.Println(strings.Repeat("%", 10))

	// 替换
	fmt.Println(strings.Replace("xyzxyzxxxxyz", "xy", "mn", -1))
	fmt.Println(strings.Replace("xyzxyzxxxxyz", "xy", "mn", 1))
	fmt.Println(strings.ReplaceAll("xyzxyzxxxxyz", "xy", "mn"))

	// 首字母大写
	fmt.Println(strings.ToTitle("my name is abcd"))
	fmt.Println(strings.ToUpper("my name is abcd"))
	fmt.Println(strings.ToLower("MY NAME IS ABCD"))

	// trim  去除
	fmt.Println(strings.Trim("abcdefabc", "bc"))
	fmt.Println(strings.Trim("abcdefabc", "abc"))
	fmt.Println(strings.TrimSpace(" \n\f\tabcdefabc\t")) // 去除空白符
	fmt.Println(strings.TrimLeft("cabcdefabc", "abc"))   // 去除左侧匹配的字符  // 左边字符出现在子字符串中则替换
	fmt.Println(strings.TrimRight("cabcdefabc", "abc"))  // 去除右侧匹配的字符

	fmt.Println(strings.TrimPrefix("cabccabcdefabca", "abc")) // 字符串当成一个整体替换
	fmt.Println(strings.TrimSuffix("cabcdefabca", "abc"))
}
