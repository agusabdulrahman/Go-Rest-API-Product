package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Completed   bool    `json:"completed"`
}

var products = []Product{
	{ID: 1, Name: "Dell XPS 13", Description: "13-inch laptop", Price: 1500.00},
	{ID: 2, Name: "MacBook Pro 16-inch", Description: "16-inch laptop", Price: 2500.00},
	{ID: 3, Name: "HP Spectre x360", Description: "13-inch 2-in-1 laptop", Price: 1200.00},
	{ID: 4, Name: "Asus ROG Zephyrus G15", Description: "15-inch gaming laptop", Price: 2000.00},
	{ID: 5, Name: "Lenovo ThinkPad X1 Carbon", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 6, Name: "Microsoft Surface Laptop 4", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 7, Name: "Acer Predator Helios 300", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 8, Name: "Microsoft Surface Book 2", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 9, Name: "Razer Blade", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 10, Name: "LG Gram", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 11, Name: "Lenovo Yoga C940", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 12, Name: "HP Envy 13", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 13, Name: "Acer Swift 5", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 14, Name: "Acer Swift 7", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 15, Name: "Microsoft Surface Laptop 4", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 16, Name: "HP Envy x360", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 17, Name: "Lenovo Yoga C740", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 18, Name: "MSI GS65 Stealth Thin", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 19, Name: "Dell G5 15n", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 20, Name: "HP Pavilion x360", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 21, Name: "Lenovo Legion 5", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 22, Name: "Acer Predator Triton 500", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 23, Name: "Asus VivoBook S15", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 24, Name: "Microsoft Surface Pro 7", Description: "14-inch business laptop", Price: 1800.00},
	{ID: 25, Name: "Razer Blade Pro 17", Description: "14-inch business laptop", Price: 1800.00},
}

func getProducts(context *gin.Context) {
	pageSize := 10 // number of products per page
	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid page number"})
		return
	}

	// calculate start and end indices of products for the requested page
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > len(products) {
		endIndex = len(products)
	}

	context.IndentedJSON(http.StatusOK, gin.H{
		"page":     page,
		"pageSize": pageSize,
		"total":    len(products),
		"products": products[startIndex:endIndex],
	})
}


func addProduct(context *gin.Context) {
	var newProduct Product

	if err := context.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)

	context.IndentedJSON(http.StatusCreated, newProduct)
}

func getProduct(context *gin.Context) {
	id := context.Param("id")
	product, err := getProductById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, product)
}

func toggleProductStatus(context *gin.Context) {
	id := context.Param("id")
	product, err := getProductById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	product.Completed = !product.Completed

	context.IndentedJSON(http.StatusOK, product)
}

func deleteProduct(context *gin.Context) {
	id := context.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid product ID"})
		return
	}
	for i, t := range products {
		if t.ID == intID {
			products = append(products[:i], products[i+1:]...)
			context.IndentedJSON(http.StatusOK, gin.H{"message": "product deleted"})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}


func getProductById(id string) (*Product, error) {
    intID, err := strconv.Atoi(id)
    if err != nil {
        return nil, errors.New("invalid product ID")
    }
    for i, t := range products {
        if t.ID == intID {
            return &products[i], nil
        }
    }
    return nil, errors.New("product not found")
}

func filterProducts(context *gin.Context) {
    var filteredProducts []Product
    query := context.Request.URL.Query()
    name := query.Get("name")
    priceMin, err := strconv.ParseFloat(query.Get("price_min"), 64)
    if err != nil {
        priceMin = 0
    }
    priceMax, err := strconv.ParseFloat(query.Get("price_max"), 64)
    if err != nil {
        priceMax = 0
    }
    completed, err := strconv.ParseBool(query.Get("completed"))
    if err != nil {
        completed = false
    }

    for _, product := range products {
        if name != "" && !strings.Contains(strings.ToLower(product.Name), strings.ToLower(name)) {
            continue
        }
        if priceMin > 0 && product.Price < priceMin {
            continue
        }
        if priceMax > 0 && product.Price > priceMax {
            continue
        }
        if completed && !product.Completed {
            continue
        }
        filteredProducts = append(filteredProducts, product)
    }

    context.IndentedJSON(http.StatusOK, filteredProducts)
}



func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProduct)
	router.PATCH("/products/:id", toggleProductStatus)
	router.POST("/products", addProduct)
	router.DELETE("/products/:id", deleteProduct)
	router.GET("/products/filter", filterProducts)

	router.Run("localhost:8080")
}
