package user

import (
	"context"
	"gozerozj/user-api/internal/svc"
	"gozerozj/user-api/internal/types"
	"gozerozj/user-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserRequest) (resp *types.AddUserResponse, err error) {
	// todo: add your logic here and delete this line
	user := &model.User{
		Tel:      req.Tel,
		NickName: req.NickName,
	}
	//插入数据库
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}

	resp = &types.AddUserResponse{
		Id: 1,
		AddUserRequest: types.AddUserRequest{
			Tel:      user.Tel,
			NickName: user.NickName,
		},
	}

	return
}
