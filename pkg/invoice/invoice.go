package invoice

import (
	"github.com/valentinolopezbruno/composicion/pkg/customer"
	"github.com/valentinolopezbruno/composicion/pkg/invoiceitem"
)

// Invoice is a struct
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	item    []invoiceitem.Item
}

func New(country, city string, client customer.Customer, item []invoiceitem.Item) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		item:    item,
	}
}

func (i *Invoice) SetTotal() {
	for _, item := range i.item {
		i.total += item.Value()
	}
}
