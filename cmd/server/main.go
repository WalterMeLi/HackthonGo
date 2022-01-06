package main

import (
	"context"
	"database/sql"
	"net/http"

	// utils "github.com/WalterMeLi/HackthonGo/utils"

	"github.com/WalterMeLi/HackthonGo/internal/product"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// 	utils.Replace("#$%#", ";", "../../datos/customers.txt", "../../datos/invoices.txt", "../../datos/products.txt", "../../datos/sales.txt")

func GetAllProducts(s product.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		p, err := s.GetAll(context.Background())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": p,
		})
	}
}

func LoadProducts(s product.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := s.LoadData(context.Background())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Datos Cargados",
		})
	}
}

func main() {
	db, _ := sql.Open("mysql", "root:@/hackthon_db")
	repoP := product.NewRepository(db)
	serviceP := product.NewService(repoP)
	router := gin.Default()

	router.GET("/products", GetAllProducts(serviceP))
	router.GET("/loadProducts", LoadProducts(serviceP))
	router.Run()
}
