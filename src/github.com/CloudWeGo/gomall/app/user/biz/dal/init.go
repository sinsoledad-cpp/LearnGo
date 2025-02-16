package dal

import (
	"github.com/CloudWeGo/gomall/app/user/biz/dal/mysql"
	"github.com/CloudWeGo/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

