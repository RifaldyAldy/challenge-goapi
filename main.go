package main

import (
	"submission-project-enigma-laundry/config"
	"submission-project-enigma-laundry/controller/customer"
	"submission-project-enigma-laundry/controller/product"
	transaksidetailscreate "submission-project-enigma-laundry/controller/transaksi_details_create"
	transaksidetailsget "submission-project-enigma-laundry/controller/transaksi_details_get"

	"github.com/gin-gonic/gin"
)

func main() {
	defer config.Dbconnect().Close()
	r := gin.Default()
	{
		products := r.Group("/products")
		products.GET("/", product.GetList)
		products.GET("/:id", product.GetId)
		products.POST("/", product.Create)
		products.PUT("/:id", product.Update)
		products.DELETE("/:id", product.Delete)
	}

	{
		customers := r.Group("/customers")
		customers.POST("/", customer.Create)
		customers.GET("/:id", customer.Get)
		customers.GET("/", customer.GetList)
		customers.PUT("/:id", customer.Update)
		customers.DELETE("/:id", customer.Delete)
	}

	{
		transactions := r.Group("/transactions")
		transactions.GET("/:id_bill", transaksidetailsget.GetTransaction)
		transactions.GET("/", transaksidetailsget.GetList)
		transactions.POST("/", transaksidetailscreate.Create)
	}

	// r.POST("/products", product.Create)
	// r.GET("/products", product.GetList)
	// r.GET("/products/:id", product.GetId)
	// r.PUT("/products/:id", product.Update)
	// r.DELETE("/products/:id", product.Delete)
	r.Run(":8080")

}
