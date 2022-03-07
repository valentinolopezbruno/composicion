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
	item    invoiceitem.Items
}

func (i *Invoice) SetTotal() { //funcion para setear valor total del slice item
	i.total = i.item.Total()
}

func New(country, city string, client customer.Customer, item invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		item:    item,
	}
}
