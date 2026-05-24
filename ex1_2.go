package main

import (
	"slices"
	"sort"
)

// Étant donné une liste de stocks,
// retourner les productIDs dont la quantité totale
// toutes magasins confondus est égale à zéro.

// Input  : [{SKU-001, 0}, {SKU-001, 0}, {SKU-002, 5}]
// Output : ["SKU-001"]

//Cas limites
// Tous les produits ont au moins 1 quantité totale
// Tous les produits ont 0 de quantité totale
// 1 seul stock en tout (Quantité == 0)
// 1 seul stock en tout (Quantité > 0)
// Liste des stocks vide

//Stock avec une quantité négative

//Complexité (O(n)):

// Retourne les productIDs dont la quantité totale
// toutes magasins confondus est égale à zéro.
//stocks : liste des stocks
//return : les productIDs dont la quantité totale
// toutes magasins confondus est égale à zéro.

func QuantZero(stocks []Stock) []string {
	qts := make(map[string]int)
	res := []string{}

	for _, s := range stocks {
		p := s.ProductID
		qts[p] = qts[p] + s.Quantity
	}

	for p, q := range qts {
		if q == 0 {
			res = append(res, p)
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	slices.SortFunc(res, func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	})

	return res
}
