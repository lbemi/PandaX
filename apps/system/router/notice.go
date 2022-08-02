package router

import (
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/gin-gonic/gin"
	"pandax/apps/system/api"
	"pandax/apps/system/services"
)

func InitNoticeRouter(router *gin.RouterGroup) {
	s := &api.NoticeApi{
		DeptApp:   services.SysDeptModelDao,
		NoticeApp: services.SysNoticeModelDao,
	}
	notice := router.Group("notice")

	notice.GET("list", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("获取通知分页列表").Handle(s.GetNoticeList)
	})

	notice.POST("", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("添加通知信息").Handle(s.InsertNotice)
	})

	notice.PUT("", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("修改通知信息").Handle(s.UpdateNotice)
	})

	notice.DELETE(":noticeId", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("删除通知信息").Handle(s.DeleteNotice)
	})
}
