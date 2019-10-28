package database

import (
	"fmt"
	"gin-project-template/pkg/log"
	"gin-project-template/pkg/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Orm *gorm.DB

func init() {
	if conf, err := util.FetchMysqlConf(); err == nil {
		//root:123456@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local&timeout=10ms
		args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=%dms", conf.Username, conf.Password, conf.Host, conf.Database, conf.Charset, conf.Timeout)
		Orm, err = gorm.Open("mysql", args)
		if err != nil {
			_ = log.Error(fmt.Sprintf("mysql connect error %v", err))
		}
		if Orm.Error != nil {
			_ = log.Error(fmt.Sprintf("databanse error %v", Orm.Error))
		}
	} else {
		_ = log.Error(fmt.Sprintf("fetch mysql config error %v", err))
	}
}
