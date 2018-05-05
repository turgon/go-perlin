package perlin

import "testing"

func TestNoise(t *testing.T) {
	p := New(0)
	t.Log(p.Noise(1.2, 1.1, 1.9))
}

func BenchmarkNoise(b *testing.B) {
	for n := 0; n < b.N; n++ {
		p := New(int64(n))
		p.Noise(1.2, 1.1, 1.9)
	}
}
