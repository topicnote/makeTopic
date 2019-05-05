//getnews → makeTopic (this file)
//トピックを生成する

package maketopic

import (
	"bufio"
	"os/exec"
	// "github.com/topicnoteteam/getnews"
)

// NewsStruct ニュース構造体
type NewsStruct struct {
	ID    uint64
	Title string
	URL   string
}

// TopicStruct トピック構造体
type TopicStruct struct {
	ID          uint64
	AddedNewsID []uint64
}

// MakeTopic トピック構造体のスライスを生成する
func MakeTopic(newsList []NewsStruct) []TopicStruct {
	topicList := []TopicStruct{}
	var nTopicID uint64
	w2v := exec.Command("python3", "w2v.py") //in:[NewsTitle string]  out:[TopicID int]
	topicIDsbyte, _ := w2v.Output()          //ニュースのtopicIDを取得
	// news毎にTopicIDを取得してtopicListに追加する
	flg := 0
	scanner := bufio.NewScanner(topicIDsbyte)

	for _, news := range newsList {
		nTopicID = scanner.ScanInts()
		flg = 0
		if len(topicList) != 0 {
			for _, topic := range topicList { //一致するTopicIDがあれば追加、なければTopicを追加
				if topic.ID == nTopicID {
					append(topic.AddedNewsID, nTopicID)
					flg = 1
					break
				}
			}
		}
		if flg == 0 {
			newTopic := TopicStruct{nTopicID, []uint64{1, news.ID}} // 新規Topic
			append(topicList, newTopic)
		}
	}

	return topicList
}
