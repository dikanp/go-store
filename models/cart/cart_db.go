package cart

import (
	"AltaStore/models"	
	"AltaStore/models/product"
)


type Cart struct {
	models.GormModel
	Products  []product.Product `gorm:"many2many:cart_products;" json:"product"`
	CustomerID 	uint	`json:"customer_id"`
}
