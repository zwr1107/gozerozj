package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozerozj/user-api/internal/logic/user"
	"gozerozj/user-api/internal/svc"
	"gozerozj/user-api/internal/types"
)

func EditUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewEditUserLogic(r.Context(), svcCtx)
		resp, err := l.EditUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
