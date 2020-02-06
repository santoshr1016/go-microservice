package app

import (
	"github.com/santoshr1016/go-microservice/mvc/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}
