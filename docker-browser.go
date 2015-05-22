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
type Image struct{
	Imagename string
}
type LBInstance struct{
	ServerIp string
	ServerPort string
}
type LBContent struct{
	Status int
	Instance LBInstance
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
func (d *DockerClient) GetIBAddr(lbaddr string,image string)(*LBContent,error){
	ld,err :=newClient("http://192.168.0.196:3000")
	if err != nil{
		return nil,err	
	}
	data := Image{
		Imagename:image,
	}
	body,_,err := ld.do("POST","/api/dispatcher/v1.0/container/create",data,true,nil)
	var li LBContent
	err = json.Unmarshal(body.&li)
	if err != nil{
		return nil,err
	}
	return &li,nil
}
func (d *DockerClient) DirectDockerRun(data RunData) (*RunRes, error){

	body,_,err := d.Docker.do("POST","/api/coderunner",data,true,nil)
	if err != nil{
		return nil,err
	}
	var li RunRes
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
