package dependency

import "golang-webapi/conf"

func GetDependencies(cfg *conf.Conf) []interface{} {
	return []interface{}{
		cfg,
		InitDB(cfg),
		InitRedis(cfg),
	}
}
