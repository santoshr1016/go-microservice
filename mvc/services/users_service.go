package services

import (
	"github.com/santoshr1016/go-microservice/mvc/domain"
	"github.com/santoshr1016/go-microservice/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUserFromDb(userId)
}
