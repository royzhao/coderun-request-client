package client

import (
	"encoding/json"
	//	"net/http"
	// "net/url"
	"fmt"
)

type DockerClient struct {
	Docker *client
}

type Code_type struct {
	Filename string
	Content  string
}
type Cmd_type struct {
	Cmd  string
	Args string
}
type RunData struct {
	Id      string
	Workdir string
	Code    Code_type
	Cmds    []Cmd_type
}

type RunRes struct {
	Status  int
	Message string
}

func NewDockerClient(endpoint string) (*DockerClient, error) {
	c, err := newClient(endpoint)
	if err != nil {
		return nil, err
	}
	return &DockerClient{
		Docker: c,
	}, nil
}

func (d *DockerClient) DirectDockerRun(data RunData) (*RunRes, error){

	body,_,err := d.Docker.do("POST","/api/coderunner",data,true,nil)
	if err != nil{
		return nil,err
	}
	var li Runres
	err = json.Unmarshal(body,&li)
	if err != nil{
		return nil,err
	}
	return &li,nil
}
func (d *DockerClient) DockerRun(data RunData, image string) (*RunRes, error) {
	body, _, err := d.Docker.do("POST", "/runner/"+image, data, true, nil)
	if err != nil {
		return nil, err
	}
	var li RunRes
	err = json.Unmarshal(body, &li)
	if err != nil {
		return nil, err
	}
	fmt.Println(li)
	return &li, nil
}
