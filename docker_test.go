package client

import (
	"testing"
	"fmt"
)

func Test_lb(t *testing.T){

	dc,err := NewDockerClient("http://baidu.com")

	runaddr,err := dc.GetIBAddr("http://192.168.0.196:3000","test-zplzpl:2")
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(runaddr)
}
