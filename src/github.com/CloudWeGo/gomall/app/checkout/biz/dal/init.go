package dal

import (
	"github.com/CloudWeGo/gomall/app/checkout/biz/dal/mysql"
	"github.com/CloudWeGo/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
