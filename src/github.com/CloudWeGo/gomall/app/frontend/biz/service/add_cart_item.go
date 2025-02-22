package service

import (
	"context"

	rpccart "github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/cart"
	cart "github.com/CloudWeGo/gomall/app/frontend/hertz_gen/frontend/cart"
	common "github.com/CloudWeGo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/CloudWeGo/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/CloudWeGo/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  uint32(req.ProductNum),
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
