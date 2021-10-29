package handle

import (
	logic "api/internal/logic/sys/admin"
	"net/http"

	"api/internal/svc"
	"api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func AdminLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAdminLoginLogic(r.Context(), ctx)
		resp, err := l.AdminLogin(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
