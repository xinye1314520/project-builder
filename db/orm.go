package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"project-builder/util"
	"os"
	"project-builder/entity"
)

var (
	dbEngine   *xorm.Engine
	checkCount int
)

func init() {
	var err error

	//主库添加
	url := util.Config.Section("mysql").Key("url").String()
	dbEngine, err = xorm.NewEngine("mysql", url)
	dbEngine.SetMaxIdleConns(10)
	dbEngine.SetMaxOpenConns(200)
	dbEngine.ShowSQL(true)
	dbEngine.ShowExecTime(true)

	if err != nil {
		fmt.Printf("Fail to connect to db: %v", err)
		os.Exit(1)
	}
	syncErr := dbEngine.Sync2(new(entity.User))
	if syncErr != nil{
		fmt.Printf("sync err:%v\n", syncErr)
	}
}

func GetDBEngine() *xorm.Engine {
	return dbEngine
}

func CheckDB() {
	checkCount ++
	fmt.Printf("Check db count : %d\n", checkCount)

	if dbEngine != nil {
		// Raw SQL
		pingErr := dbEngine.Ping()

		if pingErr != nil {
			fmt.Printf("Check db warn : %v\n", pingErr)
		} else {
			fmt.Println("Check db normal")
		}
	} else {
		fmt.Println("DB conn error")
	}
}