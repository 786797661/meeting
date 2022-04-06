package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"meeting/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewMeetingRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .实例化Data 用于数据库，redis对象的实例话
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, rdb: rdb}, cleanup, nil
}

//NewDB
func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func NewRedis(c *conf.Data) *redis.Client {
	println(c.Redis.Addr, c.Redis.Db)
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		DB:           int(c.Redis.Db),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})

	//心跳
	pong, err := rdb.Ping().Result()
	log.Info(pong, err) // Output: PONG <nil>
	println(pong, err)
	//if err := rdb.Close(); err != nil {
	//	log.Error(err)
	//	println(err)
	//}
	return rdb
}
