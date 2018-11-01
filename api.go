package main

import (
	"github.com/j-forster/Waziup-API/api"
	routing "github.com/julienschmidt/httprouter"
)

var router = routing.New()

func init() {

	router.POST("/auth/token", api.GetToken)
	router.GET("/auth/permissions", api.GetPermissions)
}
