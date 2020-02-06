package domain

import (
	"github.com/santoshr1016/go-microservice/mvc/utils"
	"net/http"
)

var(
	inMemoryUsers = map[int64]*User{
		123: &User{
			Id:        1,
			FirstName: "Santosh",
			LastName:  "Kumar",
			Email:     "san123@gmail.com",
		},
		12: &User{
			Id:        2,
			FirstName: "Shruthi",
			LastName:  "GK",
			Email:     "gk123@gmail.com",
		},
	}

)


func GetUserFromDb(userId int64) (*User, *utils.ApplicationError) {
	if user := inMemoryUsers[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    "User not found",
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
	//user := inMemoryUsers[userId]
	//if user == nil {
	//	return nil, errors.New(fmt.Sprintf("User not Found %v", userId))
	//}
	//return user, nil
}
