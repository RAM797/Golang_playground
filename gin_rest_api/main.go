package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type book struct {
	ISIN   string  `json:"isin"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBooks(c *gin.Context) {
	var newBook book
	err := c.BindJSON(&newBook)
	if err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByISIN(c *gin.Context) {
	isin := c.Param("isin")
	fmt.Println(isin)
	for _, b := range books {
		if b.ISIN == isin {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func updatePrice(c *gin.Context) {
	isin := c.Request.FormValue("isin")
	price := c.Request.FormValue("price")
	fmt.Println(isin, price)
	for i, b := range books {
		fmt.Println(b.ISIN)
		if b.ISIN == isin {
			newPrice, err := strconv.ParseFloat(price, 64)
			if err != nil {
				c.IndentedJSON(http.StatusNotModified, gin.H{"error": "incorrect price type"})
				return
			}
			b.Price = newPrice
			books[i] = b
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func removeBook(c *gin.Context) {
	isin := c.Request.FormValue("isin")
	for i, b := range books {
		if b.ISIN == isin {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, books)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

var books = []book{
	{ISIN: "1US", Title: "Harry Potter and the Chamber of Secrets", Author: "JK Rowling", Price: 45.40},
	{ISIN: "2US", Title: "The subtle art of not giving a fu*k", Author: "Mark Manson", Price: 32.90},
	{ISIN: "3US", Title: "Sapiens", Author: "Yuval Noah Harari", Price: 76.00},
	{ISIN: "1IN", Title: "Life is what you make it", Author: "Preeti Shenoy", Price: 20.00}}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:isin", getBookByISIN)
	router.POST("/books", postBooks)
	router.PUT("/books", updatePrice)
	router.DELETE("/books", removeBook)
	router.Run("localhost:8000")
}
