package client

import (
	"encoding/json"
	//	"net/http"
	// "net/url"
	// "fmt"
)

type LBClient struct {
	LB *client
}

type Image struct {
	Imagename string
}

type LBInstance struct {
	ServerIP   string
	ServerPort int
}
type LBContent struct {
	Status   int
	Instance LBInstance
}

// type RunRes struct {
// 	Status  int
// 	Message string
// }

func NewLBClient(endpoint string) (*LBClient, error) {
	c, err := newClient(endpoint)
	if err != nil {
		return nil, err
	}
	return &LBClient{
		LB: c,
	}, nil
}

func (d *LBClient) PrepareImage(image string) (*LBContent, error) {
	data := Image{
		Imagename: image,
	}
	body, _, err := d.LB.do("POST", "/api/dispatcher/v1.0/container/create", data, true, nil)
	if err != nil {
		return nil, err
	}
	var li LBContent
	err = json.Unmarshal(body, &li)
	if err != nil {
		return nil, err
	}
	return &li, nil

}
