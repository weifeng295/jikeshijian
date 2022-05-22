package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	logging "github.com/sirupsen/logrus"
	"log"
	"warp-demo/basic/config"
)

func initMysql() {

	var err error

	// 创建连接
	mysqlDB, err = gorm.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	mysqlDB.SingularTable(true)

	// 最大连接数
	mysqlDB.DB().SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	// 最大闲置数
	mysqlDB.DB().SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())
	// 激活链接
	if err = mysqlDB.DB().Ping(); err != nil {
		logging.Error(err)
	}
}
