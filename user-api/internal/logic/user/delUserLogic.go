package user

import (
	"context"

	"gozerozj/user-api/internal/svc"
	"gozerozj/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelUserLogic) DelUser(req *types.DelUserRequest) (resp *types.DelUserResponse, err error) {
	// todo: add your logic here and delete this line

	id := req.Id
	//查询数据库
	err = l.svcCtx.UserModel.Delete(l.ctx, id)
	if err != nil {
		return nil, err
	}
	resp = &types.DelUserResponse{
		Id: id,
	}
	return
}
