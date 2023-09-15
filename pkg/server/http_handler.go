package server

import (
	"fmt"
	"io"
	"net/http"
)

var _ http.Handler = (*VehiclesHTTPHandler)(nil)

type VehiclesHTTPHandler struct {
}

func NewVehiclesHTTPHandler() *VehiclesHTTPHandler {
	return &VehiclesHTTPHandler{}
}

func (v *VehiclesHTTPHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "Hi")
	if err != nil {
		fmt.Println("erro ao responder requisição", err)
	}
}
