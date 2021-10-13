package dependency

import (
	"golang-webapi/conf"
	"golang-webapi/services"
)

func GetDependencies(cfg *conf.Conf) []interface{} {
	db := InitDB(cfg)
	return []interface{}{
		cfg,
		db,
		InitRedis(cfg),

		services.NewUserService(db),
	}
}
