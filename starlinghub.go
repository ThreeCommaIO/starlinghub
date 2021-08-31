package starlinghub

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

const baseURL = "/api/connect/v1/"

type StarlingConnectService struct {
	key   string
	sling *sling.Sling
}

type Params struct {
	Key string `url:"key"`
}

func New(url string, key string, httpClient *http.Client) *StarlingConnectService {
	path := fmt.Sprintf("%s%s", url, baseURL)
	return &StarlingConnectService{
		key:   key,
		sling: sling.New().Client(httpClient).Base(path),
	}
}

func (s *StarlingConnectService) ListDevices() ([]Device, *http.Response, error) {
	listDevices := new(ListDevices)
	path := "devices"
	params := Params{Key: s.key}
	resp, err := s.sling.New().Get(path).QueryStruct(params).ReceiveSuccess(listDevices)

	return listDevices.Devices, resp, err
}

func (s *StarlingConnectService) GetDevice(id string) (Device, *http.Response, error) {
	listDevice := new(ListDevice)
	path := fmt.Sprintf("devices/%s", id)
	params := Params{Key: s.key}
	resp, err := s.sling.New().Get(path).QueryStruct(params).ReceiveSuccess(listDevice)

	return listDevice.Properties, resp, err
}
