package server

import (
	"golang.org/x/sync/singleflight"
	"staging/dao"
	"staging/pkg/settings"
)

type Server struct {
	dao    *dao.Dao
	single *singleflight.Group
}

func InitServer(app *settings.AppConfig) *Server {
	svc := &Server{
		dao:    dao.Init(app),
		single: new(singleflight.Group),
	}
	return svc
}
