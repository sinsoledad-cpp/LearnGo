package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin_qimi/bluebell/logic"
	"go.uber.org/zap"
)

// ---- 跟社区相关的 ----

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区(community_id, community_name) 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailerHandler(c *gin.Context) {
	// 获取社区id
	itStr := c.Param("id")// 从URL中获取id，获取URL参数
	id, err := strconv.ParseInt(itStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 根据id获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
