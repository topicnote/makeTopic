//getnews → makeTopic (this file)
//トピックを生成する

package main

import (
	"fmt"
	// ""
)

func getKeyWord() (title string, url string, KeyWords string) { //getnewsで取得した記事からキーワードの抽出
	return "title", "URL", "keywords"
}

func main() {
	fmt.Println("Hi")
	title, url, KeyWords := getKeyWord()
	fmt.Println(title, url, KeyWords)
}
