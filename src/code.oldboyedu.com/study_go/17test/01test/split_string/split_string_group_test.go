package split_string

import (
	"reflect"
	"testing"
)

// 测试组
func TestSplit01(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := []testCase{
		{"babcbef", "b", []string{"", "a", "c", "ef"}},
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"abcef", "bc", []string{"a", "ef"}},
		{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}
	for _, tc := range testGroup {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("want:%#v got:%#v\n", tc.want, got)
		}
	}
}

// 表格驱动测试
func TestSplit02(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case_1": {"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case_2": {"a:b:c", ":", []string{"a", "b", "c"}},
		"case_3": {"abcef", "bc", []string{"a", "ef"}},
		"case_4": {"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})

	}
}

// 性能测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

//性能比较测试
func benchmarkFib(b *testing.B,n int){
	for i:=0;i<b.N;i++{
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B){
	benchmarkFib(b,1)
}

func BenchmarkFib10(b *testing.B){
	benchmarkFib(b,10)
}