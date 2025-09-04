package model

type Product struct {
	SkuID int64
	Count uint16
	Name  string
	Price uint32
}

type Cart struct {
	UserID     int64
	Item       map[int64]*Product // ключ - ску продукта
	TotalPrice uint32
}
