package main

import (
	"testing"
)

func equalUnordered(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	counts := make(map[string]int)
	for _, v := range a {
		counts[v]++
	}
	for _, v := range b {
		counts[v]--
		if counts[v] < 0 {
			return false
		}
	}
	return true
}

func TestQuantZero(t *testing.T) {
	tests := []struct {
		name   string
		stocks []Stock
		want   []string
	}{
		{
			name: "nominal case",
			stocks: []Stock{{ProductID: "SKU-001", StoreID: "PARIS", Quantity: 10},
				{ProductID: "SKU-001", StoreID: "LYON", Quantity: 5},
				{ProductID: "SKU-002", StoreID: "PARIS", Quantity: 0},
				{ProductID: "SKU-003", StoreID: "PARIS", Quantity: 0},
				{ProductID: "SKU-002", StoreID: "BORDEAUX", Quantity: 3},
				{ProductID: "SKU-004", StoreID: "PARIS", Quantity: 0},
				{ProductID: "SKU-004", StoreID: "BORDEAUX", Quantity: 0}},
			want: []string{"SKU-003", "SKU-004"},
		},
		{
			name: "All prducts have at least 1 quantity",
			stocks: []Stock{{ProductID: "SKU-001", StoreID: "PARIS", Quantity: 10},
				{ProductID: "SKU-001", StoreID: "LYON", Quantity: 5},
				{ProductID: "SKU-002", StoreID: "PARIS", Quantity: 0},
				{ProductID: "SKU-003", StoreID: "PARIS", Quantity: 1},
				{ProductID: "SKU-002", StoreID: "BORDEAUX", Quantity: 3},
				{ProductID: "SKU-004", StoreID: "PARIS", Quantity: 0},
				{ProductID: "SKU-004", StoreID: "BORDEAUX", Quantity: 2}},
			want: []string{},
		},
		{
			name: "All products have 0 quantity",
			stocks: []Stock{{ProductID: "SKU-001", StoreID: "PARIS", Quantity: 0},
				{ProductID: "SKU-001", StoreID: "LYON", Quantity: 0},
				{ProductID: "SKU-002", StoreID: "PARIS", Quantity: 0},
				{ProductID: "SKU-003", StoreID: "PARIS", Quantity: 0},
				{ProductID: "SKU-002", StoreID: "BORDEAUX", Quantity: 0},
				{ProductID: "SKU-004", StoreID: "PARIS", Quantity: 0},
				{ProductID: "SKU-004", StoreID: "BORDEAUX", Quantity: 0}},
			want: []string{"SKU-001", "SKU-002", "SKU-003", "SKU-004"},
		},
		{
			name:   "stock lists empty",
			stocks: []Stock{},
			want:   []string{},
		},
		{
			name:   "One stock with quantity == 0",
			stocks: []Stock{{ProductID: "SKU-001", StoreID: "PARIS", Quantity: 0}},
			want:   []string{"SKU-001"},
		},
		{
			name:   "One stock with quantity > 0",
			stocks: []Stock{{ProductID: "SKU-001", StoreID: "PARIS", Quantity: 10}},
			want:   []string{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ps := QuantZero(tc.stocks)
			if !equalUnordered(ps, tc.want) {
				t.Errorf("QuantZero() res = %s, want %s", ps, tc.want)
			}
		})
	}
}
