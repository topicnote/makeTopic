package maketopic

import (
	"fmt"
	"makeTopic"
	"~/connDB"
)

func UpdateTopic(st *makeTopic.TopicStruct) {
	db := connDB.Conndb()
	defer db.Close()
	query := "INSERT INTO table (id, newsID) VALUES (" + st.ID + "," + st.AddednewsID + ") ON DUPLICATED KEY UPDATE newsID =" + st.AddednewsID
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res) //for debug
}

func InsertNews(st *makeTopic.NewsStruct) {
	db := connDB.Conndb()
	defer db.Close()
	query := "INSERT INTO table (id, title, url) VALUES (" + st.ID + "," + st.Title + "," + st.URL + ")"
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res) //for debug
}
