//getnews → makeTopic (this file)
//トピックを生成する

package maketopic

import (
	"bufio"
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
	w2v := exec.Command("python3", "w2v.py") //in:[NewsTitle string]  out:[TopicID int]
	topicIDsbyte, _ := w2v.Output()          //ニュースのtopicIDを取得
	// news毎にTopicIDを取得してtopicListに追加する
	appendTopicFlg := false
	newTopicFlg := false
	r := bytes.NewReader(topicIDsbyte)
	scanner := bufio.NewScanner(r)

	for _, news := range newsList {
		scanner.Scan()
		nTopicIDstr = scanner.Text()
		appendTopicFlg = false
		newTopicFlg = true
		if strings.Contains(nTopicIDstr, "*") {
			newTopicFlg = false
			strings.TrimRight(nTopicIDstr, "*")
		}
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
			newTopic := TopicStruct{nTopicID, []uint64{1, news.ID}, newTopicFlg} // 新規Topic
			topicList = append(topicList, newTopic)
		}
	}

	return topicList
}
