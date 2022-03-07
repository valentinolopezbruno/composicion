package main

import (
	"fmt"

	"github.com/valentinolopezbruno/composicion/pkg/customer"
	"github.com/valentinolopezbruno/composicion/pkg/invoice"
	"github.com/valentinolopezbruno/composicion/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Argentina",
		"Villa Maria",
		customer.New("Valentino Lopez", "Alver 1940", "3534223133"),
		[]invoiceitem.Item{
			invoiceitem.New(1, "Mouse Razer", 7999),
			invoiceitem.New(2, "Teclado Hyperex", 8000),
			invoiceitem.New(3, "Mouse Pad Hyperex XL", 6000),
			invoiceitem.New(4, "HeadSets Hyperex Cloud Alpha 2", 13000),
			invoiceitem.New(5, "Monitor LG 24´´ (144hz)", 20000),
		},
	)

	i.SetTotal()

	fmt.Println(" FIRST INVOICE: ", i)

}
