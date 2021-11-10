package svc

import (
	"api/internal/config"
	"api/internal/middleware"
	"api/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	AdminModel model.AdminsModel
	CheckUrl   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		AdminModel: model.NewAdminsModel(conn),
		CheckUrl:   middleware.NewCheckUrlMiddleware().Handle,
	}
}
