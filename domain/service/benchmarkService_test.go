package service

import "testing"

func Test_simpleQuadraticAlgorithm(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"test", 100000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			simpleQuadraticAlgorithm(tt.size)
		})
	}
}
