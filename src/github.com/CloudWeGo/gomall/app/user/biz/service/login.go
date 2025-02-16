package service

import (
	"context"
	"errors"

	"github.com/CloudWeGo/gomall/app/user/biz/dal/mysql"
	"github.com/CloudWeGo/gomall/app/user/biz/model"
	user "github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	row, err := model.GetByEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password)); err != nil {
		return nil, err
	}
	resp = &user.LoginResp{
		UserId: int64(row.ID),
	}
	return resp, nil
}
