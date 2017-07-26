package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/advancedlogic/GoOse"
)

var (
	url = flag.String("url",
		"http://www.gov.cn/premier/2017-07/25/content_5213316.htm",
		"news url")
)

func main() {
	flag.Parse()
	g := goose.New()
	article, err := g.ExtractFromURL(*url)
	if err != nil {
		println(err)
		return
	}
	article.RawHTML = ""
	out, _ := json.Marshal(article)
	//println(string(out))
	fmt.Println(string(out))
	/*
		println("title", article.Title)
		println("description", article.MetaDescription)
		println("keywords", article.MetaKeywords)
		println("content", article.CleanedText)
		println("url", article.FinalURL)
		println("top image", article.TopImage)
	*/
}
