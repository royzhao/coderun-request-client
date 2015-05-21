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
	user, err := sso.GetUserInfo("1", "rjLR0wIP04y7Nxybal09Re2Xn3ZeOy1Pdzt0kyXsr0IlGbsN9tVd72jC0wZPuGRJ", "act_get=get&user_id=1&user_by=user_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(user)
}
