package iteration

import testing "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// for b.Loop() {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
