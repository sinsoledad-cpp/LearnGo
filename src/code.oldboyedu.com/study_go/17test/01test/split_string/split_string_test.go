package split_string

import (
	"reflect"
	"testing"
)

// 测试
func TestSplit1(t *testing.T) {
	got := Split("babcbef", "b")
	want := []string{"", "a", "c", "ef"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v,but got:%v\n", want, got)
	}
}

func TestSplit2(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v,but got:%v\n", want, got)
	}
}

func TestSplit3(t *testing.T) {
	got := Split("abcdef", "bc")
	want := []string{"a", "def"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v,but got:%v\n", want, got)
	}
}
