package util

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

var Config *ini.File
var (
	RootPath = "D:\\gopro\\src\\project-builder"
)

func init() {
	//RootPath项目绝对路径
	var err error
	Config, err = ini.Load(RootPath + "/conf/config.ini");
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}

func GetProPath() string {
	return RootPath
}
