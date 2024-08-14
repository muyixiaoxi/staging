package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"staging/pkg/settings"
)

func initDB(m *settings.MySQLConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.DB)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		zap.L().Error("gorm init failed %v", zap.Error(err))
	}

	return db
}
