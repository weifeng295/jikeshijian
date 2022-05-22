package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"sync"
	"warp-demo/basic/config"
)

var (
	inited  bool
	mysqlDB *gorm.DB
	m       sync.RWMutex
)

// Init 初始化数据库
func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] DB 已经初始化过")
		logging.Info(err.Error())
		return
	}

	// 如果配置声明使用mysql
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}

	inited = true
}

// GetDB 获取db
func GetDB() *gorm.DB {
	mysqlDB.LogMode(true) //开启调试模式
	return mysqlDB
}
