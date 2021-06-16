package services

import (
	"bot/daos"
	"bot/redis"
	"bot/shared"
)

//Service : ""
type Service struct {
	Daos   *daos.Daos
	Shared *shared.Shared
	Redis  *redis.RedisCli
}

//IService : ""
type IService interface {
}

//GetService :""
func GetService(dao *daos.Daos, s *shared.Shared, Redis *redis.RedisCli) *Service {
	return &Service{dao, s, Redis}
}
