package main

import (
	"flag"
	"fmt"
	"github.com/thatguystone/swan"
)

var (
	url = flag.String("url",
		"http://www.gov.cn/premier/2017-07/25/content_5213316.htm",
		"news url")
)

func main() {
	flag.Parse()
	article, err := swan.FromURL(*url)
	if err != nil {
		println(err)
		return
	}
	fmt.Println(article.Meta)
}
