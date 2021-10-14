package dependency

import (
	"golang-webapi/conf"
	"golang-webapi/model"

	"github.com/kataras/golog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var loggerDB = golog.New()

func InitDB(cfg *conf.Conf) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               cfg.MySQL,
		DefaultStringSize: 256,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		msg := "open mysql err:" + err.Error()
		loggerDB.Error(msg)
		panic(msg)
	}

	//create tables
	db.AutoMigrate(model.Tables...)
	return db
}
