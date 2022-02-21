package db

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

var Mysql *gorm.DB
var RedisConn *redis.Pool
