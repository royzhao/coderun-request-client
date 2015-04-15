//client
package client

import (
	"encoding/json"
	"fmt"
	//	"net/http"
	"net/url"
)

type SSOClient struct {
	SClient *client
}

func NewSSOClient(endpoint string) (*SSOClient, error) {
	c, err := newClient(endpoint)
	if err != nil {
		return nil, err
	}
	return &SSOClient{
		SClient: c,
	}, nil
}

type Login struct {
	Is_login string `json:"is_login" yaml:"is_login"`
	Uid      string `json:"u_id,omitempty" yaml:"u_id,omitempty"`
	Uname    string `json:"u_name,omitempty" yaml:"u_name,omitempty"`
}

//func (s *SSOClient) IsLogin(path string, data interface{}) (string, error) {
func (c *SSOClient) IsLogin(method string, path string, data url.Values) (string, error) {
	body, _, err := c.SClient.do(method, path, nil, false, data)
	//	fmt.Println(status)
	if err != nil {
		return "false", err
	}
	var li Login
	err = json.Unmarshal(body, &li)
	if err != nil {
		return "false", err
	}
	return li.Is_login, nil
}

type Out struct {
	Is_out string `json:"is_logout" yaml:"is_logout"`
}

func (c *SSOClient) Logout(method string, path string, data url.Values) (string, error) {
	body, _, err := c.SClient.do(method, path, nil, false, data)
	if err != nil {
		return "true", err
	}
	var lo Out
	err = json.Unmarshal(body, &lo)
	if err != nil {
		return "true", err
	}
	return lo.Is_logout, nil
}
