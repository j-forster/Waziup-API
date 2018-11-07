package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/j-forster/Waziup-API/tools"

	router "github.com/julienschmidt/httprouter"
)

type Device struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	GatewayId string    `json:"gateway_id"`
	Sensors   []*Sensor `json:"sensors"`
}

var devices map[string]*Device

////////////////////

func GetDevices(resp http.ResponseWriter, req *http.Request, params router.Params) {

	buf := bytes.Buffer{}
	buf.Write([]byte{'['})
	for _, device := range devices {
		// NOT-CONFORM: Returns also last_value.
		data, err := json.MarshalIndent(device, "", "  ")
		if err != nil {
			http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		buf.Write(data)
		buf.Write([]byte{','})
	}
	buf.Write([]byte{']'})
	resp.Write(buf.Bytes())
}

func CreateDevice(resp http.ResponseWriter, req *http.Request, params router.Params) {

	data, err := tools.ReadAll(req.Body)
	if err != nil {
		http.Error(resp, "Request Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	device := &Device{}
	err = json.Unmarshal(data, device)
	if err != nil {
		http.Error(resp, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}
	if device.Id == "" {
		// NOT-CONFORM: Create a unique id if no id was given.
		device.Id = uuid.New().String()
	}
	devices[device.Id] = device

	// NOT-CONFORM: Return id on success.
	resp.Header().Set("Content-Type", "text/plain")
	resp.Write([]byte(device.Id))
}

func GetDevice(resp http.ResponseWriter, req *http.Request, params router.Params) {
	id := params.ByName("device_id")

	device := devices[id]
	if device == nil {
		http.Error(resp, "Not Found: Device not found.", http.StatusNotFound)
		return
	}

	data, err := json.Marshal(device)
	if err != nil {
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(data)
}
