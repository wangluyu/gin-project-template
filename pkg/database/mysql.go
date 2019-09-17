package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Orm *gorm.DB

func init() {
	var err error
	Orm, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/tallybook?charset=utf8mb4&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Orm.Error != nil {
		fmt.Printf("databanse error %v", Orm.Error)
	}
}
