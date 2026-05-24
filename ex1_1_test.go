package main

import (
	"testing"
)

func TestCountStocksOfProduct(t *testing.T) {

	tests := []struct {
		name      string
		stocks    []Stock
		wantCount int
	}{
		{
			name: "Nominal",
			stocks: []Stock{{ProductID: "SKU-001", StoreID: "PARIS", Quantity: 10},
				{ProductID: "SKU-001", StoreID: "LYON", Quantity: 5},
				{ProductID: "SKU-002", StoreID: "PARIS", Quantity: 0},
				{ProductID: "SKU-002", StoreID: "BORDEAUX", Quantity: 3}},
			wantCount: 15,
		},
		{
			name:      "Stock List Empty",
			stocks:    []Stock{},
			wantCount: 0,
		},
		{
			name: "No stock for the product",
			stocks: []Stock{{ProductID: "SKU-002", StoreID: "PARIS", Quantity: 0},
				{ProductID: "SKU-002", StoreID: "BORDEAUX", Quantity: 3}},
			wantCount: 0,
		},
		{
			name:      "Only one stock",
			stocks:    []Stock{{ProductID: "SKU-001", StoreID: "PARIS", Quantity: 10}},
			wantCount: 10,
		},
		{
			name: "Product with all stocks at 0 quantity",
			stocks: []Stock{{ProductID: "SKU-001", StoreID: "PARIS", Quantity: 0},
				{ProductID: "SKU-001", StoreID: "LYON", Quantity: 0},
				{ProductID: "SKU-002", StoreID: "PARIS", Quantity: 0},
				{ProductID: "SKU-002", StoreID: "BORDEAUX", Quantity: 3}},
			wantCount: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			count := CountStocksOfProduct(tc.stocks, "SKU-001")
			if count != tc.wantCount {
				t.Errorf("CountStocksOfProduct() Count = %d, want %d", count, tc.wantCount)
			}
		})
	}
}
