package 代理模式

import "testing"

func TestProxy(t *testing.T) {
	var sub Subject
	sub = &Proxy{}
	res := sub.Do()
	if res != "pre:real:after" {
		t.Fatal("fail")
	}
}
