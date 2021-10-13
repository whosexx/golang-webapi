package dependency

import (
	"golang-webapi/conf"
	"golang-webapi/services"

	"github.com/kataras/golog"
)

func GetDependencies(cfg *conf.Conf) []interface{} {
	db := InitDB(cfg)
	return []interface{}{
		cfg,
		db,
		InitRedis(cfg),
		golog.New(),
		services.NewUserService(db),
	}
}
