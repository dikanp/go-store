package controllers

import (
	// "AltaStore/configs"
	// "AltaStore/models/product"
	"AltaStore/configs"
	"AltaStore/models/cart"
	"AltaStore/models/cartitems"
	"net/http"

	// "net/http"
	"strconv"

	"github.com/labstack/echo"
)


func DeleteCartControllers(c echo.Context) error {
	cartId, err := strconv.Atoi(c.Param("cartId"))
	productId, _ := strconv.Atoi(c.Param("productId"))
	
	var dataCart cart.Cart
	var dataCartItems cartitems.CartItems
	
	err = configs.DB.Where("cart_id = ? AND product_id = ?", cartId, productId).Delete(&dataCartItems).Error
	
	configs.DB.Preload("Products").First(&dataCart, cartId)


	if err != nil {
		return c.JSON(http.StatusInternalServerError, cart.ResponseCart{
			false, "Failed to Delete Data", nil,
		})
	}
	
	// _ = id
	return c.JSON(http.StatusOK, dataCart)
}