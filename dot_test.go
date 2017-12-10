package dot

import (
	"gonum.org/v1/gonum/blas/blas32"
	"gonum.org/v1/gonum/blas/blas64"
	cblas "gonum.org/v1/netlib/blas/netlib"
	"math/rand"
	"testing"
)

const (
	THOUSAND = 1000
	MILLION  = 1000000
)

func getFloat32RandVector(dim int) []float32 {
	vec := make([]float32, dim)
	for i := 0; i < len(vec); i++ {
		vec[i] = float32(rand.NormFloat64())
	}
	return vec
}

func getFloat64RandVector(dim int) []float64 {
	vec := make([]float64, dim)
	for i := 0; i < len(vec); i++ {
		vec[i] = rand.NormFloat64()
	}
	return vec
}

func getBlas32RandVector(dim int) blas32.Vector {
	vec := getFloat32RandVector(dim)
	return blas32.Vector{
		Inc:  1,
		Data: vec,
	}
}

func getBlas64RandVector(dim int) blas64.Vector {
	vec := getFloat64RandVector(dim)
	return blas64.Vector{
		Inc:  1,
		Data: vec,
	}
}

func BenchmarkDotFloat32ThousandGo(b *testing.B) {
	x := getBlas32RandVector(THOUSAND)
	y := getBlas32RandVector(THOUSAND)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = blas32.Dot(THOUSAND, x, y)
	}
}

func BenchmarkDotFloat32ThousandCgo(b *testing.B) {
	x := getFloat32RandVector(THOUSAND)
	y := getFloat32RandVector(THOUSAND)
	impl := cblas.Implementation{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = impl.Sdot(THOUSAND, x, 1, y, 1)
	}
}

func BenchmarkDotFloat32MillionGo(b *testing.B) {
	x := getBlas32RandVector(MILLION)
	y := getBlas32RandVector(MILLION)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = blas32.Dot(MILLION, x, y)
	}
}

func BenchmarkDotFloat32MillionCgo(b *testing.B) {
	x := getFloat32RandVector(MILLION)
	y := getFloat32RandVector(MILLION)
	impl := cblas.Implementation{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = impl.Sdot(MILLION, x, 1, y, 1)
	}
}

func BenchmarkDotFloat64ThousandGo(b *testing.B) {
	x := getBlas64RandVector(THOUSAND)
	y := getBlas64RandVector(THOUSAND)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = blas64.Dot(THOUSAND, x, y)
	}
}

func BenchmarkDotFloat64ThousandCgo(b *testing.B) {
	x := getFloat64RandVector(THOUSAND)
	y := getFloat64RandVector(THOUSAND)
	impl := cblas.Implementation{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = impl.Ddot(THOUSAND, x, 1, y, 1)
	}
}

func BenchmarkDotFloat64MillionGo(b *testing.B) {
	x := getBlas64RandVector(MILLION)
	y := getBlas64RandVector(MILLION)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = blas64.Dot(MILLION, x, y)
	}
}

func BenchmarkDotFloat64MillionCgo(b *testing.B) {
	x := getFloat64RandVector(MILLION)
	y := getFloat64RandVector(MILLION)
	impl := cblas.Implementation{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = impl.Ddot(MILLION, x, 1, y, 1)
	}
}
