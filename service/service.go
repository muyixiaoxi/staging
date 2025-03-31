package service

import (
	"golang.org/x/sync/singleflight"
	"staging/dao"
	"staging/pkg/settings"
)

type Service struct {
	dao    *dao.Dao
	single *singleflight.Group
}

func InitServer(app *settings.AppConfig) (*Service, error) {
	dao, err := dao.Init(app)
	if err != nil {
		return nil, err
	}
	svc := &Service{
		dao:    dao,
		single: new(singleflight.Group),
	}
	return svc, nil
}
