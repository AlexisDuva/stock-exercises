package main

import (
	"image/color"
	"log"
	"math/rand"
	"slices"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

type RandomizedSet struct {
	Values []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{Values: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	for _, v := range this.Values {
		if v == val {
			return false
		}
	}
	this.Values = append(this.Values, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	for i, v := range this.Values {
		if v == val {
			this.Values = slices.Delete(this.Values, i, i+1)
			return true
		}
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.Values[rand.Int()%len(this.Values)]
}

func hist(dist []float64, name, title string) {
	n := len(dist)
	vals := make(plotter.Values, n)
	for i := 0; i < n; i++ {
		vals[i] = dist[i]
	}

	plt := plot.New()
	plt.Title.Text = title
	hist, err := plotter.NewHist(vals, 25) // 25 bins
	if err != nil {
		log.Println("Cannot plot:", err)
	}
	hist.FillColor = color.RGBA{R: 255, G: 127, B: 80, A: 255} // coral color
	plt.Add(hist)

	err = plt.Save(400, 200, name+".png")
	if err != nil {
		log.Panic(err)
	}
}
