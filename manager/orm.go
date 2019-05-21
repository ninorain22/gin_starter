package manager

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/ninorain22/gintest/config"
	"fmt"
	"os"
	"github.com/go-xorm/core"
	xrc "github.com/go-xorm/xorm-redis-cache"
	)

var Engine *xorm.Engine

func init() {
	engine, err := xorm.NewEngine(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		fmt.Printf("db connect failed %s\n", err)
		os.Exit(-1)
	}
	engine.SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	engine.SetMaxOpenConns(config.DBConfig.MaxOpenConns)
	engine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, config.DBConfig.Prefix))
	engine.ShowSQL(true)

	engine.Logger().SetLevel(core.LOG_DEBUG)

	// 使用reids做缓存
	cacher := xrc.NewRedisCacher(
		config.RedisConfig.URL,
		config.RedisConfig.Password,
		config.RedisConfig.Expire,
		engine.Logger(),
		)
	engine.SetDefaultCacher(cacher)

	Engine = engine
}