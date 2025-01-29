package svc

import (
	"rpc/internal/config"

	"github.com/fengfengzhidao/go-zero/common/init_gorm"
	"github.com/fengfengzhidao/go_zero/_04rpc_study/user_gorm/modelss"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := init_gorm.InitGorm(c.Mysql.DataSource)
	db.AutoMigrate(&modelss.UserModel{})
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
