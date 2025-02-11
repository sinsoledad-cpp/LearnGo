// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"

	"github.com/cloudwego/biz-demo/book-shop/app/user/service"
	user "github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/user"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserReq) (resp *user.CreateUserResp, err error) {
	resp = user.NewCreateUserResp()

	if len(req.GetUserName()) == 0 || len(req.GetPassword()) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserReq) (resp *user.MGetUserResp, err error) {
	resp = user.NewMGetUserResp()

	if len(req.Ids) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.Users = users
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserReq) (resp *user.CheckUserResp, err error) {
	resp = user.NewCheckUserResp()

	if len(req.Password) == 0 || len(req.UserName) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userId, err := service.NewUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = userId
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}
