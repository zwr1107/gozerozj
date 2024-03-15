package user

import (
	"context"

	"gozerozj/user-api/internal/svc"
	"gozerozj/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	//获取id
	Id := req.Id
	//查询数据库
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, Id)
	if err != nil {
		return nil, err
	}
	//返回数据
	return &types.UserInfoResponse{
		NickName: user.NickName,
		Tel:      user.Tel,
	}, nil
}
