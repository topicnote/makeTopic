//getnews → makeTopic (this file)
//トピックを生成する

package main

import (
	"fmt"

	// "github.com/topicnoteteam/getnews"
	"topicNote/getnews"
)

// func getKeyWord() (title string, url string, KeyWords string) { //getnewsで取得した記事からキーワードの抽出
func getKeyWord(article getnews.NewsStruct) (title string, url string) { //getnewsで取得した記事からキーワードの抽出
	title = article.Title
	url = article.Url
	// return "title", "URL", "keywords"
	return title, url
}

func main() {
	articles := getnews.Getnews()
	for _, article := range articles {
		title, url := getKeyWord(article)
		fmt.Println(title, url)
	}
}
