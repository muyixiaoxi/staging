package dao

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"staging/pkg/settings"
)

type Dao struct {
	db  *gorm.DB
	rdb *redis.Client
}

func Init(app *settings.AppConfig) (*Dao, error) {
	db, err := initDB(app.MySQLConfig)
	if err != nil {
		return nil, err
	}
	dao := &Dao{
		db:  db,
		rdb: initRDB(app.RedisConfig),
	}
	return dao, nil
}
