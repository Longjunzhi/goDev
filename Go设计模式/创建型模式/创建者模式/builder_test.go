package 创建者模式

import "testing"

func TestBuilder1_GetResult(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != "123" {
		t.Fatalf("Builder1 fail res: %s", res)
	}
}

func TestBuilder2_GetResult(t *testing.T) {
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != 6 {
		t.Fatalf("Builder2 fail res: %d", res)
	}
}
