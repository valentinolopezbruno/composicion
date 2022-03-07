package invoiceitem

type Item struct {
	id      uint
	product string
	value   float64
}

func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

type Items []Item

// le decimos que creamos un tipo Items que sera creado en base a un SLICE de tipo Item
// resaltar que son dos tipos diferentes

func (is Items) Total() float64 { //aca definimos una funcion del tipo Items
	var total float64
	for _, item := range is { //decimos que recorra el slice Items y que sume los valores de cara Item
		total += item.value
	}
	return total //retornamos valor total
}

func NewItems(items ...Item) Items {
	var is Items
	for _, item := range items {
		is = append(is, item)
	}
	return is
}
