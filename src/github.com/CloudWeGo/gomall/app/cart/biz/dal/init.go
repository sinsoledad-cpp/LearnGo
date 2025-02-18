package dal

import (
	"github.com/CloudWeGo/gomall/app/cart/biz/dal/mysql"
	"github.com/CloudWeGo/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
