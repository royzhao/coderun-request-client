//client
package client

import (
	"encoding/base64"
	"encoding/json"
	//	"net/http"
	"fmt"
	"net/url"
)

type UserCode struct {
	Code      string
	Key       string
	Str_alert string
}
type UserInfo struct {
	User_id         string
	User_name       string
	User_mail       string
	User_nick       string
	User_time       string
	User_time_login string
}
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

//get user info by id

func (c *SSOClient) GetUserInfo(app_id int, app_key string, args string) (UserInfo, error) {
	body, _, err := c.SClient.do("GET", "/api/api.php?"+args, nil, false, nil)
	var code UserCode
	var info UserInfo
	if err != nil {
		return info, err
	}
	err = json.Unmarshal(body, &code)
	if err != nil {
		return info, err
	}
	decode := fmt.Sprintf("?mod=code&act_get=decode&app_id=%d&app_key=%s&code=%s&key=%s", app_id, app_key, code.Code, code.Key)
	body, _, err = c.SClient.do("GET", "/api/api.php"+decode, nil, false, nil)
	if err != nil {
		return info, err
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		return info, err
	}
	info.User_id = base64.StdEncoding.DecodeString([]byte(info.User_id))
	info.User_mail = base64.StdEncoding.DecodeString([]byte(info.User_mail))
	info.User_name = base64.StdEncoding.DecodeString([]byte(info.User_name))
	info.User_nick = base64.StdEncoding.DecodeString([]byte(info.User_nick))
	info.User_time = base64.StdEncoding.DecodeString([]byte(info.User_time))
	info.User_time_login = base64.StdEncoding.DecodeString([]byte(info.User_time_login))
	return info, err
}

//func (s *SSOClient) IsLogin(path string, data interface{}) (string, error) {
func (c *SSOClient) IsLogin(data url.Values) (Login, error) {
	body, _, err := c.SClient.do("POST", "/user_identification.php", nil, false, data)
	//	fmt.Println(status)
	var li Login
	if err != nil {
		li.Is_login = "false"
		return li, err
	}
	err = json.Unmarshal(body, &li)
	if err != nil {
		li.Is_login = "false"
		return li, err
	}
	return li, nil
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
	return lo.Is_out, nil
}
