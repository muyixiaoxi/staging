package dao

import (
	"context"
	"staging/model"
)

func (dao *Dao) Test() (test model.Test, err error) {
	err = dao.db.Find(&test).Error
	dao.rdb.Get(context.Background(), "test")
	return
}
