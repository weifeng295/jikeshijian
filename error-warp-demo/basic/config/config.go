package config

import (
	"github.com/joho/godotenv"
	logging "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"sync"
)

var (
	err error
)

var (
	mysqlConfig defaultMysqlConfig
	redisConfig defaultRedisConfig
	m           sync.RWMutex
	inited      bool
)

// Init 初始化配置
func Init() {

	m.Lock()
	defer m.Unlock()

	if inited {
		logging.Info("[Init] 配置已经初始化过")
		return
	}
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	mysqlConfig.URL = os.Getenv("MYSQL_HOST")
	mysqlConfig.Enable = true
	mysqlConfig.MaxIdleConnection, _ = strconv.Atoi(os.Getenv("MYSQL_MAXIDLECONNECTION"))
	mysqlConfig.MaxOpenConnection, _ = strconv.Atoi(os.Getenv("MYSQL_MAXOPENCONNECTION"))
	redisConfig.Enabled = true
	redisConfig.Conn = os.Getenv("REDIS_URL")
	redisConfig.Password = os.Getenv("REDIS_PASSWORD")
	redisConfig.DBNum, _ = strconv.Atoi(os.Getenv("REDIS_DBNUM"))
	redisConfig.Timeout, _ = strconv.Atoi(os.Getenv("REDIS_TIMEOUT"))
	// 标记已经初始化
	inited = true
}

// GetMysqlConfig 获取mysql配置
func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

// GetRedisConfig 获取redis配置
func GetRedisConfig() (ret RedisConfig) {
	return redisConfig
}
