package main

// Étant donné une liste de stocks :
// [{ProductID: "SKU-001", StoreID: "PARIS",  Quantity: 10},
//  {ProductID: "SKU-001", StoreID: "LYON",   Quantity: 5},
//  {ProductID: "SKU-002", StoreID: "PARIS",  Quantity: 0},
//  {ProductID: "SKU-002", StoreID: "BORDEAUX", Quantity: 3}]

// Retourner la quantité totale disponible par produit :
// {"SKU-001": 15, "SKU-002": 3}

type Stock struct {
	ProductID string
	StoreID   string
	Quantity  int
}

//Cas limites:
//Liste de stocks vide
//Liste de stocks ne contient pas de stock correspondant au produit p
// Produit uniquement en stock dans un seul store

//Complexité:
//O(n)

//count = 0
//for stock in stocks
//	if stock.productID := p.ID
//		then count += stock.quantity
//  endif
//endfor

//Calcule la quantité totale disponible par produit
//stocks : liste des stocks
// p : id du produit dont on cherche la quantité
//return : quantité totale disponible du produit
func CountStocksOfProduct(stocks []Stock, p string) int {
	count := 0
	for _, s := range stocks {
		if s.ProductID == p {
			count += s.Quantity
		}
	}
	return count
}
