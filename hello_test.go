package main

import "testing"

// xxx_test.goのような名前のファイルにある必要があります。
// テスト関数はTestという単語で始まる必要があります。
// テスト関数は1つの引数のみをとります。 t *testing.T

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello,World"

	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
