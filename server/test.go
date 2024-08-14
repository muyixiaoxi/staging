package server

import "staging/model"

func (s *Server) Test() (data model.Test, err error) {
	data, err = s.dao.Test()
	return
}
