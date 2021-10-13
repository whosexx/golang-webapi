package dependency

import (
	"fmt"
	"golang-webapi/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(cfg *conf.Conf) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               cfg.MySQL,
		DefaultStringSize: 256,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		msg := "open mysql err:" + err.Error()
		fmt.Println(msg)
		panic(msg)
	}

	//create table
	//db.Migrator().CreateTable(&model.UserInfo{})
	return db
}
