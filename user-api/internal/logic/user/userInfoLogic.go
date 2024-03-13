package user

import (
	"context"
	"errors"
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
	// 获取用户信息
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	resp = &types.UserInfoResponse{
		UserId:   user.UserId,
		Nickname: user.Nickname,
	}

	return
}