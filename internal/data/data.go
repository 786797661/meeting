package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	grpcx "google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	userV1 "meeting/api/realworld/v1"
	"meeting/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewMeetingRepo, NewUserServiceClient, NewRegistrar, NewDiscovery)

// Data .
type Data struct {
	// TODO wrapped database client
	db   *gorm.DB
	rdb  *redis.Client
	user userV1.RealworldClient
}

// NewData .实例化Data 用于数据库，redis对象的实例话
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client, uc userV1.RealworldClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, rdb: rdb, user: uc}, cleanup, nil
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

// NewUserServiceClient 链接用户服务 grpc
func NewUserServiceClient(sr *conf.Service, rr registry.Discovery) userV1.RealworldClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.User.Endpoint),
		grpc.WithDiscovery(rr),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(2*time.Second),
		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)
	if err != nil {
		panic(err)
	}
	c := userV1.NewRealworldClient(conn)
	return c
}

// NewRegistrar add consul
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	println(conf.Consul.Address)
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
