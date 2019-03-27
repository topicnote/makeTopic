//getnews → makeTopic (this file)
//トピックを生成する

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	// "github.com/topicnoteteam/getnews"
)

type encorder struct {
	id string
}

//Query 構造体
type Query struct {
	AppID     string `json:"app_id"`
	RequestID string `json:"request_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	MaxNum    string `json:"max_num"`
}

func makeQuery(token string, title string, body string) []byte {
	q := Query{AppID: token, RequestID: "key", Title: title, Body: body, MaxNum: "10"}
	s, err := json.Marshal(q)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return s
}

func getKeyword(title string, desc string) {
	// doc, _ := goquery.NewDocument(url)
	// body := doc.Get
	tokenPath := "../token"
	file, err := os.Open(tokenPath)
	if err != nil {
		fmt.Println("token not found:")
		fmt.Println(err)
		// return nil
		return
	}
	defer file.Close()

	byteToken, err := ioutil.ReadAll(file) //byteTokenはstringに直すとjson形式
	token := string(byteToken)
	queryJSON := makeQuery(token, title, desc)
	goolabURL := "https://labs.goo.ne.jp/api/keyword"

	req, err := http.NewRequest(
		"POST",
		goolabURL,
		bytes.NewBuffer(queryJSON), //queryJSON, valuwだけでkeyがない説 → 空か？ request
	)

	if err != nil {
		fmt.Println("creating request failed:")
		fmt.Println(err)
		// return nil
		return
	}
	req.Header.Set("Content-Type", "application/json") //, "application/x-www-form-urlencoded"
	fmt.Println(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client Err:")
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
	// doc, _ := goquery.NewDocumentFromResponse(resp) // respをパースできる状態にした
	// keywords := doc.keywords
	return
}

func main() {
	title := "アパート殺人 室内から柄が外れた包丁 強い力で刺したか"
	desc := "26日、東京 杉並区のアパートで32歳の保育士の女性が刃物で刺されて殺害された事件で、凶器とみられる包丁は、柄の部分が外れた状態で室内から見つかっていたことが捜査関係者への取材で分かりました。警視庁はベランダから侵入した男が女性に騒がれないように強い力で刺したとみて捜査しています。"
	getKeyword(title, desc)
}
