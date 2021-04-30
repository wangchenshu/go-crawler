package model

// Product -
type Product struct {
	ID    int    `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"type:varchar(255)"`
	Price string `json:"price" gorm:"type:varchar(100)"`
	Point string `json:"point" gorm:"type:varchar(100)"`
	Pic   string `json:"pic" gorm:"longtext"`
	Link  string `json:"link" gorm:"type:varchar(255)"`
}

// Center -
type Center struct {
	ID      int    `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"type:varchar(255)"`
	Address string `json:"address" gorm:"type:varchar(255)"`
	Phone   string `json:"phone" gorm:"type:varchar(100)"`
	Date    string `json:"date" gorm:"type:varchar(100)"`
}

// ProductCommon -
type ProductCommon struct {
	Name  string
	Price string
	Point string
	Link  string
}

// CenterCommon -
type CenterCommon struct {
	Name    string
	Address string
	Phone   string
	Date    string
}
