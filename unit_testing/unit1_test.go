package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	assert.Equal(t, kubus.Volume(), volumeSeharusnya, "Perhitungan volume salah")
}

func TestHitungLuas(t *testing.T) {
	assert.Equal(t, kubus.luas(), luasSeharusnya, "Perhitungan luas salah")
}

func TestHitungKeliling(t *testing.T) {
	assert.Equal(t, kubus.keliling(), kelilingSeharusnya, "perhitungan keliling salah")
}

func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.luas()
	}
}
