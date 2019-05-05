package mapper

import "time"

type User struct {
	Id int
	Uid int64
	Name string
	Age int
	CreateTime time.Time `xorm:"create_time"`
	UpdateTime time.Time `xorm:"update_time"`
}
