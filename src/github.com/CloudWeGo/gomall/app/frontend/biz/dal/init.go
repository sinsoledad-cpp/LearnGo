package dal

import (
	"github.com/CloudWeGo/gomall/app/frontend/biz/dal/mysql"
	"github.com/CloudWeGo/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
