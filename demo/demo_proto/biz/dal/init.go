package dal

import (
	"github.com/jdw-art/tiktok_commence/demo/demo_proto/biz/dal/mysql"
	"github.com/jdw-art/tiktok_commence/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
