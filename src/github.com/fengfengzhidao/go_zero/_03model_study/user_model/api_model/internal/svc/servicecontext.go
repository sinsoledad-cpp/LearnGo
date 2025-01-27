package svc

import (
	"api_model/internal/config"

	"github.com/fengfengzhidao/go-zero/common/init_gorm"
	"github.com/fengfengzhidao/go_zero/_03model_study/user_model/models"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := init_gorm.InitGorm(c.Mysql.DataSource)
	mysqlDb.AutoMigrate(&models.UserModel{})
	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
	}
}
