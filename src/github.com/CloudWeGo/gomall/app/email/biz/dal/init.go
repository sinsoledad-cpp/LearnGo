package dal

import (
	"github.com/CloudWeGo/gomall/app/email/biz/dal/mysql"
	"github.com/CloudWeGo/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
