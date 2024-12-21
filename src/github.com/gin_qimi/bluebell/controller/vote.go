package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin_qimi/bluebell/logic"
	"github.com/gin_qimi/bluebell/models"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteController(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			// fmt.Println("err", err)
			ResponseError(c, CodeInvalidParam)
			return
		}
		errDate := removeTopStruct(errs.Translate(trans)) // 翻译并去除掉错误提示中的结构体标识
		// fmt.Println("errDate", errDate)
		ResponseErrorWithMsg(c, CodeInvalidParam, errDate)
		return
	}
	// 获取当前请求的用户的id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 具体投票的业务逻辑
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
