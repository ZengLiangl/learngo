package main

import "testing"

func TestSubStr(t *testing.T) {
	tests := []struct {
		str string
		ans int
	}{
		{"我爱中国adsads我爱中国", 7},
		{"adasdaw", 2},
		{"121123", 1},
	}

	for _, tt := range tests {

		if str := lengthNonRepeatingSubStr(tt.str); str != tt.ans {
			t.Errorf("lengthNonRepeatingSubStr(%s) got %d expected %d", tt.str, str, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "我爱中国adsads我爱中国"
	ans := 7
	for i := 0; i < b.N; i++ {
		actual := lengthNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("lengthNonRepeatingSubStr(%s) got %d expected %d", s, actual, ans)

		}
	}

}
