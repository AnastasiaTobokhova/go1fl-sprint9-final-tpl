package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-10, 0, 1, 7, 8, 1000} {
		data := generateRandomElements(size)
		wantLen := size
		if wantLen < 0 {
			wantLen = 0
		}
		if len(data) != wantLen {
			t.Fatalf("generateRandomElements(%d): length = %d, want %d", size, len(data), wantLen)
		}
		for i, value := range data {
			if value <= 0 {
				t.Errorf("generateRandomElements(%d)[%d] = %d, want positive integer", size, i, value)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	minInt := -int(^uint(0)>>1) - 1
	maxInt := int(^uint(0) >> 1)
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{42}, 42},
		{"zero", []int{0}, 0},
		{"first", []int{9, 3, 1}, 9},
		{"middle", []int{1, 9, 3}, 9},
		{"last", []int{1, 3, 9}, 9},
		{"equal", []int{5, 5, 5}, 5},
		{"negative", []int{-9, -2, -7}, -2},
		{"mixed", []int{-3, 0, 7, -1}, 7},
		{"min_int", []int{minInt}, minInt},
		{"max_int", []int{0, maxInt, minInt}, maxInt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	for _, data := range [][]int{nil, {}, {42}, {-7}, {-9, -2, -7}, {0, 0}} {
		if got, want := maxChunks(data), maximum(data); got != want {
			t.Errorf("maxChunks(%v) = %d, want %d", data, got, want)
		}
	}
	// Проверяем максимум на каждой позиции, включая границы частей и остаток.
	for _, size := range []int{7, 8, 9, 15, 16, 17, 65} {
		for peak := 0; peak < size; peak++ {
			data := make([]int, size)
			for i := range data {
				data[i] = -10
			}
			data[peak] = -1
			if got := maxChunks(data); got != -1 {
				t.Errorf("size %d, peak %d: maxChunks = %d, want -1", size, peak, got)
			}
			for i, value := range data {
				want := -10
				if i == peak {
					want = -1
				}
				if value != want {
					t.Fatalf("maxChunks changed data[%d]", i)
				}
			}
		}
	}
}
