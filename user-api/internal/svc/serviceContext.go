package svc

import (
	"gozerozj/user-api/internal/config"
	"gozerozj/user-api/internal/middleware"
	"gozerozj/user-api/model"

	"github.com/zeromicro/go-zero/rest"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UserModel
	TestMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
		UserModel:      model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
