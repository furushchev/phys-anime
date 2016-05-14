package main

import "math"

type Color struct {
	rgba [4] float64
}

func (this *Color)R() float64 {
	return this.rgba[0]
}
func (this *Color)G() float64 {
	return this.rgba[1]
}
func (this *Color)B() float64 {
	return this.rgba[2]
}
func (this *Color)A() float64 {
	return this.rgba[3]
}

func NewColorRGBA(r,g,b,a float64) *Color {
	c := &Color {
		rgba: [4]float64{r,g,b,a},
	}
	return c
}

func NewColorHSVA(h,s,v,a float64) *Color {
	r, g, b := hsv2rgb(h,s,v)
	return NewColorRGBA(r,g,b,a)
}

func NewColorHSIA(h,s,i,a float64) *Color {
	r,g,b := hsi2rgb(h,s,i)
	return NewColorRGBA(r,g,b,a)
}

func (this *Color)HSV() (h,s,v float64) {
	r := this.R()
	g := this.G()
	b := this.B()
	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)
	v = (min + max) / 2.

	diff := max - min
	if diff == 0.0 {
		return
	}
	if v < 0.5 {
		s = diff / (max + min)
	} else {
		s = diff / (2 - max - min)
	}
	r2 := (((max - r) / 6) + (diff / 2)) / diff
	g2 := (((max - g) / 6) + (diff / 2)) / diff
	b2 := (((max - b) / 6) + (diff / 2)) / diff
	switch {
	case r == max:
		h = b2 -g2
	case g == max:
		h = (1.0 / 3.0) + r2 - b2
	case b == max:
		h = (2.0 / 3.0) + g2 - r2
	}
	return
}

func (this *Color)HSI() (h,s,i float64) {
	r,g,b := this.R(), this.G(), this.B()
	i = (r + g + b) / 3.0
	if i == 0 {
		s = 0
	} else {
		min := math.Min(math.Min(r,g), b)
		s = 1 - min / i
	}
	k := r - (g + b) / 2.0
	l := r*r + g*g + b*b - r*g - r*b - g*b
	if g >= b {
		h = math.Acos(k / math.Sqrt(l))
	} else {
		h = 1.0 - math.Acos(k / math.Sqrt(l))
	}
	return
}

func hsi2rgb(h,s,i float64) (r,g,b float64){
	for 0 > h || h >= 2 * math.Pi {
		if h < 0 {
			h += 2 * math.Pi
		} else {
			h -= 2 * math.Pi
		}
	}
	switch {
	case h == 0:
		r = i + 2 * i * s
		g = i - i * s
		b = i - i * s
	case 0 < h && h < 2 * math.Pi / 3:
		tmp := math.Cos(h) / math.Cos(2 * math.Pi / 6 - h)
		r = i + i * s * tmp
		g = i + i * s * (1.0 - tmp)
		b = i - i * s
	case 2 * math.Pi / 3 < h && h < 2 * math.Pi * 2 / 3:
		tmp := math.Cos(h - 2 * math.Pi / 3) / math.Cos(math.Pi - h)
		r = i - i * s
		g = i - i * s * tmp
		b = i + i * s * (1.0 - tmp)
	case h == 2 * math.Pi * 2 / 3.0:
		r = i - i * s
		g = i - i * s
		b = i + 2 * i * s
	case 2 * math.Pi * 2 / 3.0 < h && h < 2 * math.Pi:
		tmp := math.Cos(h - 4 * math.Pi / 3.0) / math.Cos(2 * math.Pi - h)
		r = i + i * s * (1.0 - tmp)
		g = i - i * s
		b = i + i * s * tmp
	}
	return
}

func hue2rgb(v1, v2, h float64) float64 {
	if h < 0 {
		h++
	}
	if h > 1 {
		h--
	}
	switch {
	case 6 * h < 1:
		return (v1 + (v2 - v1) * 6 * h)
	case 2 * h < 1:
		return v2
	case 3 * h < 2:
		return v1 + (v2 - v1) * ((2.0 / 3.0) - h) * 6
	}
	return v1
}

func hsv2rgb(h,s,v float64) (r,g,b float64) {
	if s == 0 {
		r = v
		g = v
		b = v
		return
	}
	var v1, v2 float64
	if v < 0.5 {
		v2 = v * (1 + s)
	} else {
		v2 = (v + s) - (v * s)
	}
	v1 = 2 * v - v2
	r = hue2rgb(v1, v2 ,h + (1.0 / 3.0))
	g = hue2rgb(v1, v2, h)
	b = hue2rgb(v1, v2, h - (1.0 / 3.0))
	return
}

// utility
func RandomColor(i int) *Color {
	h := math.Mod(float64(i), 20) * 100
	return NewColorHSIA(h, 1.0, 0.8, 1.0)
}