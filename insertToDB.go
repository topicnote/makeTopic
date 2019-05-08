package maketopic

import (
	"fmt"
	"makeTopic"
	"~/connDB"
)

func UpdateTopic(st *makeTopic.TopicStruct) res int {
	db := connDB.Conndb()
	defer db.Close()
	//newsIDに関する部分は、maketopic.goの方と擦り合わせる
	// newsIDarray := 追加前のnewsIDの配列
	newsIDarray = append(newsIDarray, st.AddednewsID) //追加するnewsIDをappend
	//topicIDが存在しない→新たなtopicIDを振り、newsIDarrayを追加する topicIDが存在する→newsIDarrayを更新する
	query := "INSERT INTO topic (id, newsID) VALUES (" + st.ID + "," + newsIDarray + ") ON DUPLICATED KEY UPDATE newsID =" + st.AddednewsID
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	fmt.Println(res) //for debug
	return 0
}

func InsertNews(st *makeTopic.NewsStruct) res int {
	db := connDB.Conndb()
	defer db.Close()
	query := "INSERT INTO news (title, url) VALUES (" + st.Title + "," + st.URL + ")"
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	fmt.Println(res) //for debug
	return 0
}
