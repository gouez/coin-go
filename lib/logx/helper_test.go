package logx

import "testing"

func TestInfo(t *testing.T) {
	Info("dddd")
}
func BenchmarkStdOut(b *testing.B) {
	for n := 0; n < b.N; n++ {

	}
}
