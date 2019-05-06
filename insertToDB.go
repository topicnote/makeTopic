package maketopic

import (
	"fmt"
	"makeTopic"
	"~/connDB"
)

func UpdateTopic(st *makeTopic.TopicStruct) {
	db := connDB.Conndb()
	defer db.Close()
	//newsIDに関する部分は、maketopic.goの方と擦り合わせる
	newsIDarray = append(newsIDarray, st.AddednewsID)
	query := "INSERT INTO topic (id, newsID) VALUES (" + st.ID + "," + newsIDarray + ") ON DUPLICATED KEY UPDATE newsID =" + st.AddednewsID
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res) //for debug
}

func InsertNews(st *makeTopic.NewsStruct) {
	db := connDB.Conndb()
	defer db.Close()
	query := "INSERT INTO news (title, url) VALUES (" + st.Title + "," + st.URL + ")"
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res) //for debug
}
