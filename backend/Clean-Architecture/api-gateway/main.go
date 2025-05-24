package main

import (
	"api-gateway/internal/inventory"
	"api-gateway/internal/order"
	"log"
	"net/http"
	"strconv"

	inventorypb "github.com/elliaaan/proto-gen/pb/inventory/github.com/elliaaan/proto-gen/pb/inventory"
	orderpb "github.com/elliaaan/proto-gen/pb/order/github.com/elliaaan/proto-gen/pb/order"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	invClient := inventory.NewInventoryClient("localhost:8080")
	orderClient := order.NewOrderClient("localhost:8081")

	// ============ INVENTORY ============

	r.POST("/products", func(c *gin.Context) {
		var req inventorypb.Product
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := invClient.CreateProduct(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, res)
	})

	r.GET("/products/:id", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		res, err := invClient.GetProductByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	r.GET("/products", func(c *gin.Context) {
		res, err := invClient.ListProducts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	r.PUT("/products/:id", func(c *gin.Context) {
		var req inventorypb.Product
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Id, _ = strconv.ParseUint(c.Param("id"), 10, 64)
		res, err := invClient.UpdateProduct(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	r.DELETE("/products/:id", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		_, err := invClient.DeleteProduct(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	})

	// ============ ORDER ============

	r.POST("/orders", func(c *gin.Context) {
		var req orderpb.Order
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := orderClient.CreateOrder(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, res)
	})

	r.GET("/orders/:id", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		res, err := orderClient.GetOrderByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	r.GET("/orders", func(c *gin.Context) {
		res, err := orderClient.ListOrders()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	r.PUT("/orders/:id", func(c *gin.Context) {
		var req orderpb.Order
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Id, _ = strconv.ParseUint(c.Param("id"), 10, 64)
		res, err := orderClient.UpdateOrder(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	r.DELETE("/orders/:id", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		_, err := orderClient.DeleteOrder(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	})

	log.Println("API Gateway running on port 8090...")
	r.Run(":8090")
}
