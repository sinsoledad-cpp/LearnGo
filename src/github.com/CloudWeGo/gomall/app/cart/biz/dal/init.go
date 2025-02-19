package dal

import (
	"github.com/CloudWeGo/gomall/app/cart/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
