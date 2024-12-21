package controller

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin_qimi/bluebell/dao/mysql"
	"github.com/gin_qimi/bluebell/logic"
	"github.com/gin_qimi/bluebell/models"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 这里采用前端 POST 发送json格式数据
	// 1. 获取参数和参数校验
	// ShouldBindJSON只能校验数据格式（字段类型，json格式）
	// var p models.ParamSignUp
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误,直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// c.JSON(http.StatusOK, gin.H{
			// 	// "msg": "请求参数有误",
			// 	"msg": err.Error(),
			// })
			ResponseError(c, CodeInvalidParam)
			return
		}
		// c.JSON(http.StatusOK, gin.H{
		// 	// "msg": "请求参数有误",
		// 	"msg": removeTopStruct(errs.Translate(trans)), //翻译错误信息
		// })
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// fmt.Println(p)

	// //手动对请求参数进行详细的业务逻辑校验
	// if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	// 	// 请求参数有误,直接返回响应
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "请求参数有误",
	// 	})
	// 	return
	// }

	//使用validator库进行参数校验
	// 2. 业务处理
	if err := logic.SignUp(p); err != nil { // 服务内部错误
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": "注册失败",
		// })
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg": "success",
	// })
}

func LoginHandler(c *gin.Context) {
	// 获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误,直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			// c.JSON(http.StatusOK, gin.H{
			// 	// "msg": "请求参数有误",
			// 	"msg": err.Error(),
			// })
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		// c.JSON(http.StatusOK, gin.H{
		// 	// "msg": "请求参数有误",
		// 	"msg": removeTopStruct(errs.Translate(trans)), //翻译错误信息
		// })
		return
	}
	// 业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			// c.JSON(http.StatusOK, gin.H{
			// 	"msg": "用户不存在",
			// })
			return
		}
		ResponseError(c, CodeInvalidPassword)
		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": "用户名或者密码错误",
		// })
		return
	}
	// 返回响应
	ResponseSuccess(c, gin.H{
		// "user_id":   user.UserID, //id值大干1<<53-1	Int64类型的最大值是1<<63-1
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg": "登录成功",
	// })
}
