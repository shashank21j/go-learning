package main

type Item interface {
	GetPrice() float64
}

type Gravel struct {
	lbs        float64
	pricePerLb float64
	grade      string
}

func (g Gravel) GetPrice() float64 {
	return g.lbs * g.pricePerLb
}

type Shovel struct {
	size  string
	price float64
}

func (s Shovel) GetPrice() float64 {
	return s.price
}

func CalculateOrderPrice(order []Item) float64 {
	total := 0.0
	for _, item := range order {
		total += item.GetPrice()
	}
	return total
}
