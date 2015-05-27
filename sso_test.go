package client

import (
	"log"
	"testing"
)

func Test_getuserinfo(t *testing.T) {
	sso, err := NewSSOClient("http://local.learn4me.com:4321")
	if err != nil {
		t.Error(err)
	}
	user, err := sso.GetUserInfo("1", "rnyA5jzYvDShxP5QR0R3Ip3gcEDcOEJ12Jwb9JLCbLjNbdetTwlcad83iAVxFen7", "act_get=get&user_id=1&user_by=user_id")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(user)
}
