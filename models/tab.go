package models

import (
	"path/filepath"
	"fmt"
	"strings"
)

type Tab struct {
	Id 		string   `json:"id"`
    Name    string   `json:"name"`
}




func getAllTabData() ([]Tab, int64) {
	var Tabs []Tab

	localpath, err := filepath.Abs("./routers/db_tb_tabs.csv")
	if err == nil {
		data, _ := readFile(localpath)
		fmt.Println(data)

		datas := strings.Split(data, "\n")

		for _, v := range datas {
			row := strings.Split(v, ",")
			//Videos[string(k)] = &Video{row[0], row[1], row[2], }

			// 地址
			//remoteUrls := strings.Split(row[7], ";")
			//var RemoteUrls []RemoteUrl
			//for _, rv := range remoteUrls {
			//	tagAndUrl := strings.Split(rv, "#")
			//	if len(tagAndUrl) > 1 {
			//		RemoteUrls = append(RemoteUrls, RemoteUrl{tagAndUrl[0], tagAndUrl[1]})
			//	} else {
			//		RemoteUrls = append(RemoteUrls, RemoteUrl{tagAndUrl[0], ""})
			//	}
			//}


			Tabs = append(Tabs,
				Tab{row[0],
					row[1],
				})
		}
	}

	count := int64(len(Tabs))

	return Tabs, count
}



func GetAllTab(args ...int64) ([]Tab, int64) {

	var Tabs []Tab

	var pageSize int64
	if (len(args) > 0) {
		pageSize = args[0]
	}


	Tabs, count := getAllTabData()


	if count < pageSize {
		return Tabs, int64(len(Tabs))
	}
	return Tabs[:pageSize], count
}