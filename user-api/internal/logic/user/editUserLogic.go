package user

import (
	"context"

	"gozerozj/user-api/internal/svc"
	"gozerozj/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserLogic {
	return &EditUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditUserLogic) EditUser(req *types.EditUserRequest) (resp *types.EditUserResponse, err error) {
	// todo: add your logic here and delete this line

	id := req.Id
	//查询数据库
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
	if err != nil {
		return nil, err
	}
	//更新数据
	user.NickName = req.NickName
	user.Tel = req.Tel
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return nil, err
	}
	//返回数据
	return &types.EditUserResponse{
		Id: user.Id,
		EditUserRequest: types.EditUserRequest{
			NickName: user.NickName,
			Tel:      user.Tel,
		},
	}, nil

}
