package controllers

import (
	"github.com/astaxie/beego"
	"smark_api/models"
	"strconv"
	"fmt"
)

type TabController struct {
	beego.Controller
}


//// @Title GetVideo
//// @Description find video by videoid
//// @Success 200 {video} models.Video
//// @Failure 403 :videoId is empty
//// @router /:videoId [get]
//func (o *VideoController) GetTab() {
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
func (o *TabController) GetAllTab() {

	// 分页大小
	pageSize := o.GetString("pageSize")
	_pageSize, _ := strconv.ParseInt(pageSize, 10, 64)

	fmt.Println("分页大小", _pageSize)
	if _pageSize <= 0 || _pageSize >= 10 {
		_pageSize = 10
	}

	vds, count := models.GetAllTab(_pageSize)

	result := make(map[string]interface{})
	result["data"] = vds
	result["count"] = count

	o.Data["json"] = result
	o.ServeJSON()
}
