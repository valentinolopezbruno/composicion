package customer

type Customer struct {
	name   string
	adress string
	phone  string
}

func New(name, adress, phone string) Customer {
	return Customer{name, adress, phone}
}
