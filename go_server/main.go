package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "mysql:3306"
	cfg.DBName = "classicmodels"

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// fmt.Printf("Products found: %v\n", products)

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/products/", getProducts)

	router.Run("0.0.0.0:8000")

}

type Product struct {
	Name     string `json:"name"`
	Year     string `json:"year"`
	Quantity int    `json:"quantity"`
}

func getProducts(c *gin.Context) {
	products, err := topTenProducts(2003)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, products)
}

// albumsByArtist queries for albums that have the specified artist name.
func topTenProducts(year int) ([]Product, error) {
	// An albums slice to hold data from returned rows.
	var products []Product

	rows, err := db.Query(`SELECT productName, YEAR(o.orderDate) as year,COUNT(od.quantityOrdered) as quantity 
	FROM products p, orderdetails od, orders o WHERE p.productCode = od.productCode AND
    od.orderNumber = o.orderNumber
	GROUP BY p.productCode, year
	HAVING (year = ?)
	ORDER BY year,quantity DESC
	LIMIT 10`, year)
	if err != nil {
		return nil, fmt.Errorf("Top ten product for year %d: %v", year, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var prod Product
		if err := rows.Scan(&prod.Name, &prod.Year, &prod.Quantity); err != nil {
			return nil, fmt.Errorf("Top ten product for year %d: %v", year, err)
		}
		products = append(products, prod)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Top ten product for year %d: %v", year, err)
	}
	return products, nil
}
