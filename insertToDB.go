package maketopic

import (
	"fmt"
	"strconv"
	"context"

	"../connDB"
	"../structs"
)


func UpdateTopic(topicList []structs.TopicStruct) (res int) {
	db := connDB.Conndb()
	defer db.Close()
	
	for index := 0; index < len(topicList); index = index + 1 {
		fmt.Println(index)
		var ctx context.Context
		var newsIDArrayString string
		if topicList[index].IsNewTopic == false {

			//着目TopicIDに所属するnewsIDの配列をDBから取得
			query := "SELECT newsid FROM topic WHERE id=" + strconv.FormatUint(topicList[index].ID, 10)
			fmt.Println(query)
			rows, err := db.QueryContext(ctx, query)
			if err != nil {
				fmt.Println("DB Exec Error", err)
				fmt.Println(query)
				return -1
			}
			if err := rows.Scan(&newsIDArrayString); err != nil {
				fmt.Println("UpdateTopic(): mySQL SELECT query 実行時エラー", err)
				return -1
			}
			
			//newsIDarray(original)に新しくきたnewsIDを加え
			for j := 0; j < len(topicList[index].AddedNewsID); j = j + 1{
				newsIDArrayString = newsIDArrayString + strconv.FormatUint(topicList[index].AddedNewsID[j], 10) + ", "
			}

			//変更後のnewsIDarrayをDBへ投げる
			query = "UPDATE topic SET newsid =\"" + newsIDArrayString + "\" WHERE id = " + strconv.FormatUint(topicList[index].ID, 10)
			_, err = db.Exec(query)
			if err != nil {
				fmt.Println("UpdateTopic(): mySQL UPDATE query 実行時エラー:", err)
				fmt.Println(query)
				return -1
			}

		} else { //isNewTopic == true
			
			for j := 0; j < len(topicList[index].AddedNewsID); j = j + 1{
				newsIDArrayString = newsIDArrayString + strconv.FormatUint(topicList[index].AddedNewsID[j], 10) +", "
			}
			query := "INSERT INTO topic (id, newsid) VALUES (" + strconv.FormatUint(topicList[index].ID, 10) + ",\"" + newsIDArrayString + "\")"
			_, err := db.Exec(query)
			if err != nil {
				fmt.Println("UpdateTopic(): mySQL INSERT query 実行時エラー:", err)
				fmt.Println(query)
				return -1
			}
		}
	}
	return 0
}

func InsertNews(newsList []structs.NewsStruct) ([]structs.NewsStruct, error) {
	db := connDB.Conndb()
	defer db.Close()
	var int64index int64
	for index := 0; index < len(newsList); index = index + 1 {
		query := "INSERT INTO news (title, url) VALUES (\"" + newsList[index].Title + "\",\"" + newsList[index].URL + "\")"
		res, err := db.Exec(query)
		if err != nil {
			fmt.Println("InsertNews(): mySQL INSERT query 実行時エラー:", err)
			fmt.Println(query)
			return nil, err
		}
		int64index, err = res.LastInsertId()
		newsList[index].ID = uint64(int64index)
	}
	return newsList, nil
}
