package dal

import (
	"github.com/jdw-art/tiktok_commence/app/frontend/biz/dal/mysql"
	"github.com/jdw-art/tiktok_commence/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
