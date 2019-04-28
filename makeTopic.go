//getnews → makeTopic (this file)
//トピックを生成する

package maketopic

import (
	"encoding/binary"
	"io"
	"log"
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

func makeTopic(newsList []NewsStruct) []TopicStruct {
	topicList := []TopicStruct{}
	w2v := exec.Command("python3", "w2v.py") //in:[NewsTitle string]  out:[TopicID int]

	// news毎にTopicIDを取得してtopicListに追加する
	for _, news := range newsList {
		stdin, err := w2v.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(stdin, news.Title)
		stdin.Close()
		out, _ := w2v.Output() //ニュースのtopicIDを取得
		nTopicID := binary.BigEndian.Uint64(out)
		flg := 0
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
