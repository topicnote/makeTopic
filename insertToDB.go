package maketopic

import (
	"fmt"
	"makeTopic"
	"~/connDB"
)

func InsertNewTopic(st *makeTopic.TopicStruct) {
	db := connDB.Conndb()
	var query string = "SELECT * FROM user" //tmp
	res, err := db.Exec(query)
	fmt.Println(query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func UpdateExistingTopic(st *makeTopic.TopicStruct) {

}

func InsertNewContent(st *makeTopic.NewsStruct) {

}
