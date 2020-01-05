package model

// Products -
type Products struct {
	ID    int    `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"type:varchar(255)"`
	Price string `json:"price" gorm:"type:varchar(100)"`
	Point string `json:"point" gorm:"type:varchar(100)"`
	Pic   string `json:"pic" gorm:"longtext"`
	Link  string `json:"link" gorm:"type:varchar(255)"`
}
