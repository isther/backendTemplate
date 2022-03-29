package controllers

import (
	"github.com/isther/web-base/conf"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error

	// 连接数据库
	DB, err = gorm.Open(postgres.Open(conf.Server.DSN()), &gorm.Config{})
	if err != nil {
		logrus.Fatalln(err)
	}

	// 绑定模型
	err = DB.AutoMigrate()
	if err != nil {
		logrus.Fatalln(err)
	}
}
