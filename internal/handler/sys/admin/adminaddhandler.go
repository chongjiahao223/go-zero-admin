package handle

import (
	"net/http"

	"api/internal/logic/sys/admin"
	"api/internal/svc"
	"api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func AdminAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddAdminReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAdminAddLogic(r.Context(), ctx)
		resp, err := l.AdminAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
