package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"pledge-backend/config"
	"pledge-backend/log"
	"time"
)

// InitRedis 初始化Redis
func InitRedis() *redis.Pool {
	log.Logger.Info("Init Redis")
	redisConf := config.Config.Redis
	// 建立连接池
	RedisConn = &redis.Pool{
		MaxIdle:     10,   // 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   0,    // 最大的激活连接数，表示同时最多有N个连接   0 表示无穷大
		Wait:        true, // 如果连接数不足则阻塞等待
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisConf.Address, redisConf.Port))
			if err != nil {
				return nil, err
			}
			// 验证密码
			c.Do("auth", redisConf.Password)
			// 选择db
			c.Do("select", redisConf.Db)
			return c, nil
		},
	}
	err := RedisConn.Get().Err()
	if err != nil {
		panic("redis init err " + err.Error())
	}
	return RedisConn
}

// 设置key、value、time
func RedisSet(key string, data interface{}, time interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("set", key, value)
	if err != nil {
		return err
	}
	if time != nil {
		_, err = conn.Do("expire", key, time.(int))
		if err != nil {
			return err
		}
	}
	return nil
}

// 获取Key
func RedisGet(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.Bytes(conn.Do("get", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// 删除Key
func RedisDelete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("del", key))
}

// 获取Hash类型
func RedisGetHash(key string) (map[string]string, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.StringMap(conn.Do("hgetall", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// 获取Heah其中一个值
func RedisGetHashOne(key, name string) (interface{}, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := conn.Do("hgetall", key, name)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// 设置Hash
func RedisSetHash(key string, data map[string]string, time interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()
	for k, v := range data {
		err := conn.Send("hset", key, k, v)
		if err != nil {
			return err
		}
	}
	err := conn.Flush()
	if err != nil {
		return err
	}

	if time != nil {
		_, err = conn.Do("expire", key, time.(int))
		if err != nil {
			return err
		}
	}
	return nil
}

// 删除Hash
func RedisDelHash(key string) (bool, error) {

	return true, nil
}

// 检查Key是否存在
func RedisExistsHash(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	exists, err := redis.Bool(conn.Do("hexists", key))
	if err != nil {
		return false
	}
	return exists
}

// 检查Key是否存在
func RedisExists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	exists, err := redis.Bool(conn.Do("exists", key))
	if err != nil {
		return false
	}
	return exists
}

// 获取Key剩余时间
func RedisGetTTL(key string) int64 {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.Int64(conn.Do("ttl", key))
	if err != nil {
		return 0
	}
	return reply
}

// set 集合
func RedisSAdd(k, v string) int64 {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := conn.Do("SAdd", k, v)
	if err != nil {
		return -1
	}
	return reply.(int64)
}

//获取集合元素
func RedisSmembers(k string) ([]string, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.Strings(conn.Do("smembers", k))
	if err != nil {
		return []string{}, errors.New("读取set错误")
	}
	return reply, err
}

type RedisEncryptionTask struct {
	RecordOrderFlowId int32  `json:"recordOrderFlow"` //密码转账表ID
	Encryption        string `json:"encryption"`      //密码串
	EndTime           int64  `json:"endTime"`         //失效截止时间
}

// 列表右侧添加数据
func RedisListRpush(listName string, encryption string) error {
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("rpush", listName, encryption)
	return err
}

// 取出列表中所有元素
func RedisListLRange(listName string) ([]string, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	res, err := redis.Strings(conn.Do("lrange", listName, 0, -1))
	return res, err
}

// 删除列表中指定元素
func RedisListLRem(listName string, encryption string) error {
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("lrem", listName, 1, encryption)
	return err
}

// 列表长度
func RedisListLength(listName string) (interface{}, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	len, err := conn.Do("llen", listName)
	return len, err
}

// list 删除整个列表
func RedisDelList(setName string) error {
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("del", setName)
	return err
}
