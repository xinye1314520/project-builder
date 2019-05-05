package service

import (
	"project-builder/db"
	"project-builder/mapper"
	"fmt"
	"project-builder/log"
	"go.uber.org/zap"
	"strconv"
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
		log.WebInfo.Info("user==", zap.String("key",strconv.Itoa(k)), zap.Any("val", v))
	}
}
