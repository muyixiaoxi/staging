package dao

import (
	"fmt"
	"staging/model"
)

func (d *Dao) Get() {
	user := &model.User{}
	d.db.Find(user)
	fmt.Println(user)
}
