package middleware

import (
	"api/internal/common/errorx"
	"encoding/json"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

type CheckUrlMiddleware struct {
}

func NewCheckUrlMiddleware() *CheckUrlMiddleware {
	return &CheckUrlMiddleware{}
}

func (m *CheckUrlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 判断请求header中是否携带了x-user-id
		userId := r.Context().Value("adminId").(json.Number).String()
		if userId == "" {
			logx.Errorf("缺少必要参数x-admin-id")
			httpx.Error(w, errorx.NewDefaultError("缺少必要参数x-admin-id"))
			return
		}

		// 判断用户是否存在

		// Passthrough to next handler if need
		next(w, r)
	}
}
