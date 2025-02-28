package dal

import (
	"github.com/CloudWeGo/gomall/app/product/biz/dal/mysql"
	"github.com/CloudWeGo/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
