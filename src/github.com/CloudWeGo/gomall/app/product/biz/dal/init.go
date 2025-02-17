package dal

import (
	"github.com/CloudWeGo/gomall/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
