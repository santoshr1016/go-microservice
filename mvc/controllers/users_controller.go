package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/santoshr1016/go-microservice/mvc/services"
	"github.com/santoshr1016/go-microservice/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request){
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil{
		userErr := &utils.ApplicationError{
			Message:    "user id is not number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		jsonValue, _ := json.Marshal(userErr)
		resp.WriteHeader(userErr.StatusCode)
		resp.Write(jsonValue)
		return

		//_, _ = resp.Write([]byte("user id must be number"))
	}
	//fmt.Println("Processing User: %v", userId)
	//log.Printf("Processing user: %v", userId)
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		// Handle the Error
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		_, _ = resp.Write([]byte(jsonValue))
		return
	}
	fmt.Println(user)
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}