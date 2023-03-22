package service

import (
	"sync"
)

type userService struct {
}

var instance *userService
var once sync.Once

// 实现单例
func GetUserService() *userService {
	once.Do(func() {
		instance = &userService{}
	})
	return instance
}
