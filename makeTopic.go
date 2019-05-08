//getnews → makeTopic (this file)
//トピックを生成する

package maketopic

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	getnews "../getNews"
)

// TopicStruct トピック構造体
type TopicStruct struct {
	ID          uint64
	AddedNewsID []uint64
	IsNewTopic  bool
}

// MakeTopic トピック構造体のスライスを生成する
func MakeTopic(newsList []getnews.NewsStruct) []TopicStruct {
	topicList := []TopicStruct{IsNewTopic: false}
	var nTopicIDstr str
	var nTopicID uint64
	w2v := exec.Command("python3", "w2v.py") //in:[NewsTitle string]  out:[TopicID int]
	topicIDsbyte, _ := w2v.Output()          //ニュースのtopicIDを取得
	// news毎にTopicIDを取得してtopicListに追加する
	appendTopicFlg := false
	newTopicFlg := false
	scanner := bufio.NewScanner(topicIDsbyte)

	for _, news := range newsList {
		nTopicIDstr = scanner.Scan().Text()
		appendTopicFlg = false
		newTopicFlg = true
		if strings.Contains(nTopicIDstr, "*") {
			newTopicFlg = false
			strings.TrimRight(nTopicIDstr)
		}
		nTopicID, err := strconv.ParseUint(nTopicIDstr, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		if len(topicList) != 0 {
			for _, topic := range topicList { //一致するTopicIDがあれば追加、なければTopicを追加
				if topic.ID == nTopicID {
					append(topic.AddedNewsID, nTopicID)
					appendTopicFlg = true
					break
				}
			}
		}
		if appendTopicFlg == false {
			newTopic := TopicStruct{nTopicID, []uint64{1, news.ID}, newTopicFlg} // 新規Topic
			append(topicList, newTopic)
		}
	}

	return topicList
}
