package main

// Product ...
type Product struct {
	ID    int64
	Name  string
	Price int64
}

func (c Product) getType() string {
	return "product"
}

// Coupan ...
type Coupan struct {
	ID       int64
	Code     string
	Discount int64
}

func (c Coupan) getType() string {
	return "coupan"
}

// Cart ...
type Cart struct {
	Products []Product
	Coupans  []Coupan
}

// CartItem ...
type CartItem interface {
	getType() string
}

func (c *Cart) add(ci CartItem) {
	switch cond := ci.getType(); cond {
	case "product":
		p := ci.(Product)
		c.Products = append(c.Products, p)
	case "coupan":
		break
	default:
		break
	}
}
