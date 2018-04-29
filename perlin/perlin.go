package perlin

import (
	"math/rand"
)

type Perlin struct {
	P []int
}

func New(seed int64) *Perlin {
	r := rand.New(rand.NewSource(seed))

	p := Perlin{
		P: make([]int, 0, 512),
	}

	p.P = r.Perm(256)
	p.P = append(p.P, p.P...)

	return &p
}

func (p *Perlin) Noise(x, y, z float64) float64 {
	X := int(x) & 255
	Y := int(y) & 255
	Z := int(z) & 255

	x = x - float64(int(x))
	y = y - float64(int(y))
	z = z - float64(int(z))

	u := fade(x)
	v := fade(y)
	w := fade(z)

	A := p.P[X] + Y
	AA := p.P[A] + Z
	AB := p.P[A+1] + Z

	B := p.P[X+1] + Y
	BA := p.P[B] + Z
	BB := p.P[B+1] + Z

	v1 := grad(p.P[AA], x, y, z)
	v2 := grad(p.P[BA], x-1, y, z)
	v3 := grad(p.P[AB], x, y-1, z)
	v4 := grad(p.P[BB], x-1, y-1, z)
	v5 := grad(p.P[AA+1], x, y, z-1)
	v6 := grad(p.P[BA+1], x-1, y, z-1)
	v7 := grad(p.P[AB+1], x, y-1, z-1)
	v8 := grad(p.P[BB+1], x-1, y-1, z-1)

	return lerp(w,
		lerp(v,
			lerp(u, v1, v2),
			lerp(u, v3, v4)),
		lerp(v,
			lerp(u, v5, v6),
			lerp(u, v7, v8)))
}

func fade(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}

func lerp(t, a, b float64) float64 {
	return a + t*(b-a)
}

func grad(hash int, x, y, z float64) float64 {
	h := hash & 15

	var u, v, r float64

	if h < 8 {
		u = x
	} else {
		u = y
	}

	if h < 4 {
		v = y
	} else {
		if h == 12 || h == 14 {
			v = x
		} else {
			v = z
		}
	}

	if h&1 == 0 {
		r = u
	} else {
		r = -u
	}

	if h&2 == 0 {
		r = r + v
	} else {
		r = r - v
	}

	return r
}
