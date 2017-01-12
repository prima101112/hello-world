package product

type Product struct {
	Id          int64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Review      []Review `json:"review"`
	Shop        Shop `json:"shop"`
}

type Review struct {
	Id     int64 `json:"id"`
	Review string `json:"review"`
}

type Shop struct {
	Id              int64 `json:"id"`
	ShopName        string `json:"shopname"`
	ShopDescription string `json:"shopdescription"`
}
