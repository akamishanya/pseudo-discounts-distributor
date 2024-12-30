package entity

type Product struct {
	Id            int64
	MarketplaceId int64
	Name          string
	IsAdultOnly   bool
	Link          string
	ImageId       int64
}
