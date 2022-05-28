package controllers

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/isther/backendTemplate/conf"
	"github.com/isther/backendTemplate/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
)

func init() {
	initSQL()
	initRedis()
	flushRedis()
}

func initSQL() {
	var err error

	// 连接数据库
	DB, err = gorm.Open(postgres.Open(conf.Server.DSN()), &gorm.Config{})
	if err != nil {
		logrus.Fatalln(err)
	}

	// 绑定模型
	err = DB.AutoMigrate(model.Model{})
	if err != nil {
		logrus.Fatalln(err)
	}
}

func initRedis() {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	Redis = redis.NewClient(conf.Server.RD())
	_, err = Redis.Ping(ctx).Result()
	if err != nil {
		logrus.Fatalln("Redis: ", err)
	}
}

func flushRedis() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	Redis.FlushAll(ctx)
}
