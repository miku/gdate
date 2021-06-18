package gdate

import "testing"

func TestAdd(t *testing.T) {
	// Table Driven Test
	var cases = []struct {
		a, b   int
		result int
	}{
		{0, 0, 0},
		{1, 1, 2},
	}
	for _, c := range cases {
		result := Add(c.a, c.b)
		if result != c.result {
			t.Fatalf("got %v, want %v", result, c.result)
		}
	}

	// result := Add(1, 1)
	// if result != 2 {
	// 	t.Fatalf("expected 2, got %v", result)
	// }
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 100)
	}
}

func BenchmarkAddSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddSlow(100, 100)
	}
}
