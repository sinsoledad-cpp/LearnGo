package dal

import (
	"github.com/CloudWeGo/gomall/app/order/biz/dal/mysql"
	"github.com/CloudWeGo/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
