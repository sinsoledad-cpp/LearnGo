package logic

import (
	"context"
	"errors"

	"rpc/internal/svc"
	"rpc/types/user"

	"github.com/fengfengzhidao/go_zero/_04rpc_study/user_gorm/modelss"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	var model modelss.UserModel
	err := l.svcCtx.DB.Take(&model, in.UserId).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return &user.UserInfoResponse{
		UserId:   uint32(model.ID),
		Username: model.Username,
	}, nil
}
