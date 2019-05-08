package maketopic

import (
	"fmt"

	"../connDB"
	"../structs"
)

func UpdateTopic(topicList []structs.TopicStruct) (res int) {
	db := connDB.Conndb()
	defer db.Close()

	// newsIDarray := 追加前のnewsIDの配列
	for index := 0; index < len(topicList); index = index + 1 {
		if topicList[index].IsNewTopic == false {

			query := "SELECT newsID FROM topic WHERE id=" + strconf.Itoa(topicList[index].ID)
			res, err := db.Exec(query)
			if err != nil {
				fmt.Println(err)
				return -1
			}
			fmt.Println(res) //for debug

			var newsIDarray = []uint64{0, 1, 2}                             //res of db.exec
			newsIDarray = append(newsIDarray, topicList[index].AddednewsID) //追加するnewsIDをappend
			//topicIDが存在しない→新たなtopicIDを振り、newsIDarrayを追加する topicIDが存在する→newsIDarrayを更新する
			str := fmt.Sprintf("%v", newsIDarray)
			query = "UPDATE topic SET newsID =" + str + " WHERE id = " + strconf.Itoa(topicList[index].ID)
			res, err = db.Exec(query)
			if err != nil {
				fmt.Println(err)
				return -1
			}
			fmt.Println(res) //for debug

		} else { //isNewTopic == true

			var newsIDarray = []int64{topicList[index].AddednewsID}
			query := "INSERT INTO topic (id, newsID) VALUES (" + topicList.ID + "," + newsIDarray + ")"
			res, err := db.Exec(query)
			if err != nil {
				fmt.Println(err)
				return -1
			}
			fmt.Println(res) //for debug
		}
	}
	return 0
}

func InsertNews(newsList []structs.NewsStruct) ([]structs.NewsStruct, error) {

	db := connDB.Conndb()
	defer db.Close()
	for index := 0; index < len(&newsList); index = index + 1 {
		query := "INSERT INTO news (title, url) VALUES (" + newsList[index].Title + "," + newsList[index].URL + ")"
		res, err := db.Exec(query)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		fmt.Println(res)           //for debug
		newsList[index].ID = index //かり
	}
	return newsList, nil
}
