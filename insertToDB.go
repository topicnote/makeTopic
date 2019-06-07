package maketopic

import (
	"fmt"
	"strcov"
	"database/sql"

	"../connDB"
	"../structs"
)


func UpdateTopic(topicList []structs.TopicStruct) (res int) {
	db := connDB.Conndb()
	defer db.Close()
	
	for index := 0; index < len(topicList); index = index + 1 {
		fmt.Printf("%v/%v -> ",index,len(topicList))
		var newsIDArrayString string

			//着目TopicIDに所属するnewsIDの配列をDBから取得
			query := "SELECT newsid FROM topic WHERE id=" + strconv.FormatUint(topicList[index].ID, 10)
			// fmt.Println(query)
			rows, err := db.Query(query)
			if err != nil {
				fmt.Println("dbExecErr")
			}
			defer rows.Close()
			rows.Next() 
			if err:=rows.Scan(&newsIDArrayString); err !=nil{
				switch {
					case err == sql.ErrNoRows:
						fmt.Println("newTopic"+strconv.FormatUint(topicList[index].ID, 10))
					case err != nil:
						fmt.Println("UpdateTopic(): mySQL SELECT query 実行時エラー", err)
						return -1
					default:
						
				}
			}

			//newsIDarray(original)に新しくきたnewsIDを加え
			for j := 0; j < len(topicList[index].AddedNewsID); j = j + 1{
				newsIDArrayString = newsIDArrayString + strconv.FormatUint(topicList[index].AddedNewsID[j], 10) + ", "
			}

			//変更後のnewsIDarrayをDBへ投げる
			// query = "UPDATE topic SET newsid =\"" + newsIDArrayString + "\" WHERE id = " + strconv.FormatUint(topicList[index].ID, 10)
			query = "INSERT INTO topic (id, newsid) VALUES (" + strconv.FormatUint(topicList[index].ID, 10) + ", \"" + newsIDArrayString + "\") ON DUPLICATE KEY UPDATE newsid=\"" + newsIDArrayString + "\";"
			_, err = db.Exec(query)
			if err != nil {
				fmt.Println("UpdateTopic(): mySQL UPDATE query 実行時エラー:", err)
				fmt.Println(query)
				return -1
			}

			fmt.Printf("DB操作成功: TopicID %v - news {%v}.\n",strconv.FormatUint(topicList[index].ID, 10),newsIDArrayString)

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
