package latihan_interface

import (
	"math"
)

type Hitung interface {
	Luas() float64
	Keliling() float64
}

type Lingkaran struct {
	diameter float64
}

type Persegi struct {
	sisi float64
}

func (l Lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * math.Pow(l.JariJari(), 2)
}

func (l Lingkaran) Keliling() float64 {
	return math.Pi * l.diameter
}

func (p Persegi) Luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p Persegi) Keliling() float64 {
	return p.sisi * 4
}
