package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func main() {
	q := "goquery"
	url := "https://pkg.go.dev/search?q=" + q

	// 发起http请求获取响应并创建Document 结构体指针
	document, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// 在document中通过选择器去查找元素
	// <tagname>
	// 标签选择题
	// 获取所有a标签
	selection := document.Find("a")

	selection.Each(func(index int, tag *goquery.Selection) {
		href, exists := tag.Attr("href")
		//tag.Html()

		fmt.Println(tag.Text(), href, exists)
	})

	fmt.Println("===========clasee===========")
	// calss 选择器
	// .className
	// 在classname下获取所有的超链接
	document.Find(".go-Content SearchResults").Find("a").Each(func(index int, tag *goquery.Selection) {
		href, exists := tag.Attr("href")
		//tag.Html()

		fmt.Println(tag.Text(), href, exists)
	})

	// id 选择器
	// #id
	fmt.Println(document.Find("#jump-to-modal").Attr("class"))
	fmt.Println(document.Find("#jump-to-modal").Html())
	fmt.Println(document.Find("#jump-to-modal").Text())

	// 复合选择器
	// tag + class
	// <div></div><div class="nav"></div><span class="nav"></div>
}
