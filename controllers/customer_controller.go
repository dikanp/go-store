package controllers

import (
	"AltaStore/configs"
	"AltaStore/models/customer"
	"net/http"

	"github.com/labstack/echo"
)

func GetCustomerController(e echo.Context) error {
	var dataCustomers []customer.Customer

	err := configs.DB.Find(&dataCustomers).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, customer.ResponseCustomer{
			false, "Failed get data Customers", nil,
		})
	}

	return e.JSON(http.StatusOK, customer.ResponseCustomer{
		true, "Success get Data Customer", dataCustomers,
	})
}

func RegisterController(e echo.Context) error {
	var customerRegister customer.CustomerRegister
	e.Bind(&customerRegister)

	var customerDB customer.Customer
	customerDB.Name = customerRegister.Name
	customerDB.Email = customerRegister.Email
	customerDB.Password = customerRegister.Password

	err := configs.DB.Create(&customerDB)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, customer.ResponseCustomer{
			false, "Failed register", nil,
		})
	}

	return e.JSON(http.StatusOK, customer.ResponseCustomerSingle{
		true, "Success registered", customerDB,
	})
}
