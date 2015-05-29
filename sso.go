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
	User_ip         string
	Str_alert       string
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

func convertBase642String(src string) string {
	des, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(des)
}

//get user info by id

func (c *SSOClient) GetUserInfo(app_id string, app_key string, args string) (UserInfo, error) {
	query := fmt.Sprintf("mod=user&%s&app_id=%s&app_key=%s", args, app_id, app_key)
	body, _, err := c.SClient.do("GET", "/api/api.php?"+query, nil, false, nil)
	var code UserCode
	var info UserInfo
	if err != nil {
		fmt.Println(body)
		return info, err
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		return info, err
	}
	if info.Str_alert != "y010102" {
		return info, newError(1, []byte("no such user,error code:"+code.Str_alert))
	}
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
