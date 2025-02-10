package dal

import (
	"github.com/CloudWeGo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/CloudWeGo/gomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
