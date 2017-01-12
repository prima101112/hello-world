package product

type Product struct {
	Id          int64
	Name        string
	Description string
	Review      []Review
	Shop        Shop
}

type Review struct {
	Id     int64
	Review string
}

type Shop struct {
	Id              int64
	ShopName        string
	ShopDescription string
}
