package models

import (
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"errors"
	"strconv"
)

type Video struct {
	Id		string   `json:"id"`
	Name	string   `json:"name"`
	Thumbnail   string   `json:"thumbnail"`
	Classify	Classify	`json:"classify"`
	Introduce   string	 `json:"introduce"`
	RemoteUrl	[]RemoteUrl 	 `json:"remote_url"`
	Timestamp   string 		`json:"timestamp"`
	Latest		string		`json:"latest"`   // 最新一集
	GeneratedAt	string		`json:"generated_at"`  // 最新一集的更新时间
	Source 		string		`json:"source"`
}

var (
	//Videos []Video
)

func init()  {
	//Videos = make(map[string]*Video)
	//Videos["1"] = &Video{"1", "aaa", "111"}
	//Videos["2"] = &Video{"2", "aaa", "111"}
}

func getAllData() ([]Video, int64) {
	var Videos []Video

	localpath, err := filepath.Abs("./routers/db_tb_videos.csv")
	if err == nil {
		data, _ := readFile(localpath)
		fmt.Println(data)

		datas := strings.Split(data, "\n")

		for k, v := range datas {

			// 第一行是说明
			if k == 0 {
				continue;
			}

			row := strings.Split(v, ",")
			//Videos[string(k)] = &Video{row[0], row[1], row[2], }

			// 地址
			remoteUrls := strings.Split(row[7], ";")
			var RemoteUrls []RemoteUrl
			for _, rv := range remoteUrls {
				tagAndUrl := strings.Split(rv, "#")
				if len(tagAndUrl) > 1 {
					RemoteUrls = append(RemoteUrls, RemoteUrl{tagAndUrl[0], tagAndUrl[1]})
				} else {
					RemoteUrls = append(RemoteUrls, RemoteUrl{tagAndUrl[0], ""})
				}
			}


			Videos = append(Videos,
				Video{row[0],
				row[1],
				row[2],
				Classify{"", row[5]},
				row[6],
				RemoteUrls,
				"1234567890",
				row[3],
				row[4],
				"mm",
			})
		}
	}

	count := int64(len(Videos))

	return Videos, count
}


func GetOneVideo(VideoId string) (video *Video, err error) {
	//if v, ok := Videos[VideoId]; ok {
	//	//	return v, nil
	//	//}

	Videos, _ := getAllData()

	id, err := strconv.ParseInt(VideoId, 10, 64)
	if v := &Videos[int(id)-1]; v != nil {
		return v, nil
	}

	return nil, errors.New("ObjectId Not Exist")
}



func GetAllVideo(args ...int64) ([]Video, int64) {

	var Videos []Video

	var pageSize int64
	if (len(args) > 0) {
		pageSize = args[0]
	}


	Videos, count := getAllData()


	if count < pageSize {
		return Videos, int64(len(Videos))
	}
	return Videos[:pageSize], count
}





func readFile(path string) (str string, err error) {
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件失败")
		fmt.Println(err)
	}

	defer fi.Close()

	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		fmt.Println("读取文件失败")
		fmt.Println(err)
	}

	str = string(fd)
	return str, err
}