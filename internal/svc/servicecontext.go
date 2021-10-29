package svc

import (
	"api/internal/config"
	"api/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	AdminModel model.AdminsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		AdminModel: model.NewAdminsModel(conn),
	}
}
