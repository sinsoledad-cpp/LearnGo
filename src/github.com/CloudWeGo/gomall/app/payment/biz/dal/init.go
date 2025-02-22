package dal

import (
	"github.com/CloudWeGo/gomall/app/payment/biz/dal/mysql"
	"github.com/CloudWeGo/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
