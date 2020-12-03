package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {

	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// ベンチマークテスト
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

// 練習問題
// strings パッケージをご覧ください。
// 便利だと思われる関数を見つけて、ここにあるようなテストを作成して実験してください。
func TestContains(t *testing.T) {

	contains := strings.Contains("hoge", "o")
	expected := true

	if contains != expected {
		t.Errorf("expected %t but got %t", contains, expected)
	}

}
