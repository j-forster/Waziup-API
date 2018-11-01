package api

import (
	"fmt"
	"net/http"

	router "github.com/julienschmidt/httprouter"
)

func GetToken(resp http.ResponseWriter, req *http.Request, params router.Params) {
	// TODO: implement
	fmt.Fprint(resp, "GetToken()")
}

func GetPermissions(resp http.ResponseWriter, req *http.Request, params router.Params) {
	// TODO: implement
	fmt.Fprint(resp, "GetPermissions()")
}
