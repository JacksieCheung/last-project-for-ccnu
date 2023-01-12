package model

type CommodityModel struct {
	Id             uint32  `json:"id" gorm:"column:id;not null" binding:"required"`
	Name           string  `json:"name" gorm:"column:name" binding:"required"`
	Info           string  `json:"info" gorm:"column:info" binding:"required"`
	Image          string  `json:"image" gorm:"column:image" binding:"required"`
	Url            string  `json:"url" gorm:"column:url" binding:"required"`
	VagueSellCount uint32  `json:"vague_sell_count" gorm:"column:vague_sell_count" binding:"required"`
	ShopId         uint32  `json:"shop_id" gorm:"column:shop_id" binding:"required"`
	CategoryId     uint32  `json:"category_id" gorm:"column:category_id" binding:"required"`
	Price          float64 `json:"price" gorm:"column:price; type:decimal(10,2)" binding:"required"`
}

func (c *CommodityModel) TableName() string {
	return "commodities"
}

// ListCommodities ... 查找 100 条记录
func ListCommodities(page, limit int) ([]*CommodityModel, error) {
	offset := (page - 1) * limit

	commodityList := make([]*CommodityModel, 0)

	query := DB.Self.Table("commodities").
		Limit(limit).
		Offset(offset)

	if err := query.Scan(&commodityList).Error; err != nil {
		return nil, err
	}

	return commodityList, nil
}
