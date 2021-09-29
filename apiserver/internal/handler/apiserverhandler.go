package handler

import (
	"net/http"

	"apiserver/internal/logic"
	"apiserver/internal/svc"
	"apiserver/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ApiserverHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewApiserverLogic(r.Context(), ctx)
		resp, err := l.Apiserver(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
