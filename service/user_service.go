package service

import (
	"project-builder/db"
	"project-builder/mapper"
	"fmt"
)

func InsertUser(user *mapper.User) (int64, error) {
	orm := db.GetDBEngine()
	row, insertErr := orm.Insert(user)
	return row, insertErr
}

func QueryUserList() {
	orm := db.GetDBEngine()
	resultMap, err := orm.Query("SELECT * FROM user")
	if err != nil {
		fmt.Printf("query err:%v", err)
	}
	for k, v := range resultMap {
		fmt.Printf("k:%v, v:%v\n", k, v)
	}
}
