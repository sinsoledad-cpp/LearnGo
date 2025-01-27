package logic

import (
	"context"
	"errors"

	"api_model/internal/svc"
	"api_model/internal/types"

	"github.com/fengfengzhidao/go_zero/_03model_study/user_model/models"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	var user models.UserModel
	err = l.svcCtx.DB.Take(&user, "username = ? and password = ?", req.Username, req.Password).Error
	if err != nil {
		return "", errors.New("登录失败")
	}
	return user.Username, nil
}
