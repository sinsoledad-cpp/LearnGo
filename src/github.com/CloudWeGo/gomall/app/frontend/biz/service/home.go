package service

import (
	"context"

	common "github.com/CloudWeGo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/CloudWeGo/gomall/app/frontend/infra/rpc"
	"github.com/CloudWeGo/gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	// fmt.Println("home service")

	// var resp = make(map[string]any)
	// items := []map[string]any{
	// 	{"Name": "T-shirt-1", "Price": "100", "Picture": "/static/image/t-shirt-1.jpg"},
	// 	{"Name": "T-shirt-2", "Price": "110", "Picture": "/static/image/t-shirt-1.jpg"},
	// 	{"Name": "T-shirt-3", "Price": "120", "Picture": "/static/image/t-shirt-1.jpg"},
	// 	{"Name": "T-shirt-4", "Price": "130", "Picture": "/static/image/t-shirt-1.jpg"},
	// 	{"Name": "T-shirt-5", "Price": "140", "Picture": "/static/image/t-shirt-1.jpg"},
	// 	{"Name": "T-shirt-6", "Price": "150", "Picture": "/static/image/t-shirt-1.jpg"},
	// }
	// resp["Title"] = "Hot Sales"
	// resp["Items"] = items
	// return resp, nil

	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"Title": "Hot sale",
		"Items": products.Products,
	}, nil
}
