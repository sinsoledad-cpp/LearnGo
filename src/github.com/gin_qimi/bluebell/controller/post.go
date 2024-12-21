package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin_qimi/bluebell/logic"
	"github.com/gin_qimi/bluebell/models"
	"go.uber.org/zap"
)

// CreatePostHandler 创建帖子的处理函数
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数的校验
	// c.ShouldBindJSON() // validator --> binding tag
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从 c 取到当前发请求的用户的ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2. 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// CreatePostHandler 获取帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数(从url中获取帖子的id)
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 根据id取出帖子数据（查数据库）
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)
}
func GetPostListHandler(c *gin.Context) {
	//获取分页参数
	page, size := getPageInfo(c)
	//获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回xiang'yin
	ResponseSuccess(c, data)
}

// GetPostListHandler2 升级版帖子列表接口
// 根据前端传来的参数动态的获取帖子列表
// 按创建时间排序 或者 按照 分数排序
// 1. 获取参数
// 2. 去redis查询id列表
// 3. 根据id去数据库查询帖子详细信息

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
	//获取分页参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//获取数据
	data, err := logic.GetPostListNew(p) //更新:合二为一

	// //获取数据
	// data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

// // GetCommunityPostListHandler 根据社区去查询帖子列表
// func GetCommunityPostListHandler(c *gin.Context) {
// 	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
// 	//获取分页参数
// 	p := &models.ParamCommunityPostList{
// 		ParamPostList: &models.ParamPostList{
// 			Page:  1,
// 			Size:  10,
// 			Order: models.OrderTime,
// 		},
// 	}
// 	if err := c.ShouldBindQuery(p); err != nil {
// 		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err))
// 		ResponseError(c, CodeInvalidParam)
// 		return
// 	}

// 	//获取数据
// 	data, err := logic.GetCommunityPostList(p)
// 	if err != nil {
// 		zap.L().Error("logic.GetCommunityPostList2() failed", zap.Error(err))
// 		ResponseError(c, CodeServerBusy)
// 		return
// 	}
// 	//返回响应
// 	ResponseSuccess(c, data)
// 	return
// }
