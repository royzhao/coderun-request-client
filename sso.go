//client
package client

import (
	"encoding/json"
	"fmt"
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
	Uid      int    `json:"u_id,omitempty" yaml:"u_id,omitempty"`
}

func (s *SSOClient) IsLogin(path string, data interface{}) (string, error) {
	body, status, err := s.SClient.do("POST", path, data, true)
	fmt.Println(status)
	if err != nil {
		fmt.Println("1")
		return "false", err
	}
	var lo Login
	err = json.Unmarshal(body, &lo)
	if err != nil {
		fmt.Println("2")
		return "false", err
	}
	return lo.Is_login, nil
}

//func (s *SSOClient) Logout(path string) error {
//	return error
//}
