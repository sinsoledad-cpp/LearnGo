package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Response(r *http.Request, w http.ResponseWriter, res any, err error) {
	if err != nil {
		body := Body{
			Code: 10086,
			Data: nil,
			Msg:  "err",
		}
		httpx.WriteJson(w, http.StatusOK, body)
		return
	}

	body := Body{
		Code: 0,
		Data: res,
		Msg:  "success",
	}
	httpx.WriteJson(w, http.StatusOK, body)
}
