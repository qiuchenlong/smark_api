package controllers

import (
	"github.com/astaxie/beego"
	"smark_api/models"
	"strconv"
	"fmt"
)

type IndexController struct {
	beego.Controller
}


//// @Title GetVideo
//// @Description find video by videoid
//// @Success 200 {video} models.Video
//// @Failure 403 :videoId is empty
//// @router /:videoId [get]
//func (o *VideoIndexController) GetVideo() {
//	videoId := o.Ctx.Input.Param(":videoId")
//	if videoId != "" {
//		ob, err := models.GetOneVideo(videoId)
//
//		if err != nil {
//			o.Data["json"] = err.Error()
//		} else {
//			result := make(map[string]interface{})
//			result["data"] = ob
//
//			o.Data["json"] = result
//		}
//	}
//	o.ServeJSON()
//}


// @Title GetAllVideo
// @Description get all videos
// @Success 200 {video} models.Video
// @Failure 403 :videoId is empty
// @router / [get]
func (o *IndexController) GetAllVideo() {

	// 分页大小
	pageSize := o.GetString("pageSize")
	_pageSize, _ := strconv.ParseInt(pageSize, 10, 64)

	fmt.Println("分页大小", _pageSize)
	if _pageSize <= 0 || _pageSize >= 10 {
		_pageSize = 10
	}

	vds, _ := models.GetAllVideo(_pageSize)

	result := make(map[string]interface{})


	action := make(map[string]interface{})
	action["data"] = "Nl7FJ976sg"
	action["type"] = "alipay_readpack"

	banner := make(map[string]interface{})
	banner["action"] = action
	banner["image"] = "https://dd.shmy.tech/static/ads/alipay/alipay_redpack.png"
	banner["name"] = "你好，Flutter"

	var b []interface{}
	b = append(b, banner);


	payload := make(map[string]interface{})
	payload["hots"] = vds
	payload["banner"] = b



	result["id"] = "1"
	result["name"] = "推荐"
	result["payload"] = payload

	o.Data["json"] = result
	o.ServeJSON()
}
