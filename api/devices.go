package api

import (
	"fmt"
	"net/http"

	router "github.com/julienschmidt/httprouter"
)

func GetDevices(resp http.ResponseWriter, req *http.Request, params router.Params) {
	// TODO: implement
	fmt.Fprint(resp, "GetDevices()")
}

func CreateDevice(resp http.ResponseWriter, req *http.Request, params router.Params) {
	// TODO: implement
	fmt.Fprint(resp, "CreateDevice()")
}

func GetDevice(resp http.ResponseWriter, req *http.Request, params router.Params) {
  id := params.ByName("device_id")
	// TODO: implement
	fmt.Fprint(resp, "GetDevice(%s)", id)
}
