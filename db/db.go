package dbLib

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"minijwc-kefu/lib"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(lib.Config.MysqlDsnLocal), &gorm.Config{PrepareStmt: true}) //Logger: logger.Discard,
	if err != nil {
		db, err = gorm.Open(mysql.Open(lib.Config.MysqlDsn), &gorm.Config{PrepareStmt: true}) //Logger: logger.Discard,
		if err != nil {
			logrus.Fatalln("数据库初始化错误", err)
		}
	}
	logrus.Info("数据库初始化完毕")
	return db
}

func NewRedis() *redis.Client {
	logrus.Info("redis地址", lib.Config.RedisAddr)
	rdb := redis.NewClient(&redis.Options{
		Addr:     lib.Config.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//检查redis链接
	_, err := rdb.Get(context.Background(), "key").Result()
	if err != nil && err != redis.Nil {
		logrus.Fatalln(err)
	}
	return rdb
}
