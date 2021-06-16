package handlers

import (
	"bot/redis"
	"bot/services"
	"bot/shared"
)

//Handler : ""
type Handler struct {
	Service *services.Service
	Shared  *shared.Shared
	Redis   *redis.RedisCli
	//PageService *pageservice.PageService
}

//GetHandler :""
func GetHandler(service *services.Service, s *shared.Shared, Redis *redis.RedisCli) *Handler {
	return &Handler{service, s, Redis}
}
