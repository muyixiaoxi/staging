package service

import (
	"staging/model"
)

func (s *Service) Test() (data model.User, err error) {
	s.dao.Get()
	return
}
