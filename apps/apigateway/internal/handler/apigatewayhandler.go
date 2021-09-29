package handler

import (
	"net/http"

	"apigateway/internal/logic"
	"apigateway/internal/svc"
	"apigateway/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ApigatewayHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewApigatewayLogic(r.Context(), ctx)
		resp, err := l.Apigateway(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
