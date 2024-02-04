package maps

import "testing"

func TestJoin(t *testing.T) {
	tests := []struct {
		values   []string
		separate string
		expected string
	}{
		{[]string{"apple", "orange", "banana"}, " ", "apple orange banana"},
		{[]string{"one", "two", "three"}, ",", "one,two,three"},
		{[]string{"a", "b", "c"}, ".", "a.b.c"},
	}

	for _, test := range tests {
		result := strings{}.Join(test.values, test.separate)
		if result != test.expected {
			t.Errorf("Join(%v, %s) = %s; expected %s", test.values, test.separate, result, test.expected)
		}
	}
}
