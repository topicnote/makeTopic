//getnews → makeTopic (this file)
//トピックを生成する

package maketopic

import (
	"bufio"
	"fmt"
	"os"
	"bytes"
	"os/exec"
	"strconv"
	"strings"

	. "../structs"
)

// MakeTopic トピック構造体のスライスを生成する
func MakeTopic(newsList []NewsStruct) []TopicStruct {
	var topicList []TopicStruct
	var nTopicIDstr string
	var nTopicID uint64
	w2vPath := os.Getenv("GOPATH") + "/src/topicNote/makeTopic/w2v.py"
	topicIDsbyte, err:= exec.Command("python3", w2vPath).Output()//ニュースのtopicIDを取得
	if err != nil {
		fmt.Println("exec err")
		return nil
	}
	// news毎にTopicIDを取得してtopicListに追加する
	appendTopicFlg := false
	r := bytes.NewReader(topicIDsbyte)
	scanner := bufio.NewScanner(r)

	for _, news := range newsList {
		scanner.Scan()
		nTopicIDstr = scanner.Text()
		fmt.Println("nTopicIDstr",nTopicIDstr)
		appendTopicFlg = false
		// if strings.Contains(nTopicIDstr, "*") {
		// 	nTopicIDstr = strings.TrimRight(nTopicIDstr, "*")
		// 	fmt.Println(nTopicIDstr)
		// }
		nTopicID, _ = strconv.ParseUint(nTopicIDstr, 10, 64)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		if len(topicList) != 0 {
			for _, topic := range topicList { //一致するTopicIDがあれば追加、なければTopicを追加
				if topic.ID == nTopicID {
					topic.AddedNewsID = append(topic.AddedNewsID, news.ID)
					appendTopicFlg = true
					break
				}
			}
		}
		if appendTopicFlg == false {
			newTopic := TopicStruct{nTopicID, []uint64{news.ID}} // 新規Topic
			topicList = append(topicList, newTopic)
		}
	}
	for index, topic := range topicList {
		fmt.Println(index)
		fmt.Println(topic.IsNewTopic)
	}
	return topicList
}
