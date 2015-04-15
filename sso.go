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

type test struct {
	ID int
}

func (s *SSOClient) IsLogin() error {
	body, _, err := s.SClient.do("GET", "/dockerapi/test", nil, false)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var t test
	err = json.Unmarshal(body, &t)
	fmt.Println(t.ID)
	return err
}

//func main() {
//	endpoint := "http://127.0.0.1:9000"
//	c, err := NewSSOClient(endpoint)
//	c.IsLogin()
//	fmt.Println(err)
//	fmt.Println("Hello World!")
//}
