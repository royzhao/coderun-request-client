package client

import (
	"encoding/json"
	//	"net/http"
	"net/url"
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
	Args []string
}
type RunData struct {
	Id      string
	Workdir string
	Code    Code_type
	Cmds    []Cmd_type
}

type RunRes struct {
	Code    string
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

func (d *DockerClient) dockerRun(data RunData) (*RunRes, error) {
	body, _, err := d.Docker.do("POST", "/api/coderunnrt", data, true, nil)
	if err != nil {
		return "false", err
	}
	var li RunRes
	err = json.Unmarshal(body, &li)
	if err != nil {
		return nil, err
	}
	return &li, nil
}
