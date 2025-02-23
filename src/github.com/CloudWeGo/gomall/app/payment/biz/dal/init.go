package dal

import (
	"github.com/CloudWeGo/gomall/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
