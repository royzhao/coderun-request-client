package client

import (
	"log"
	"testing"
)

func Test_getuserinfo(t *testing.T) {
	sso, err := NewSSOClient("http://sso.learn4me.com")
	if err != nil {
		t.Error(err)
	}
	user, err := sso.GetUserInfo("1", "rnyA5jzYvDShxP5QR0R3Ip3gcEDcOEJ12Jwb9JLCbLjNbdetTwlcad83iAVxFen7", "act_get=get&user_id=1&user_by=user_id")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(user)
}

func Test_prepareImage(t *testing.T) {
	lb, err := NewLBClient("http://192.168.0.196:3000")
	if err != nil {
		t.Error(err)
	}
	res, err := lb.PrepareImage("admin-ubuntu:1")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(res)
}

func Test_prepareImagefailed(t *testing.T) {
	lb, err := NewLBClient("http://192.168.0.196:3000")
	if err != nil {
		t.Error(err)
	}
	res, err := lb.PrepareImage("dasdas")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(res)
}
