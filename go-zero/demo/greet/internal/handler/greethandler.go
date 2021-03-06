package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goweb/go-zero/demo/greet/internal/logic"
	"goweb/go-zero/demo/greet/internal/svc"
	"goweb/go-zero/demo/greet/internal/types"
)

func GreetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGreetLogic(r.Context(), ctx)
		resp, err := l.Greet(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
