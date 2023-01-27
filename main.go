package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type billionaire struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Country  string  `json: "country"`
	NetWorth float64 ` json: "netWorth"`
	Image    string  `json:"imageUrl"`
}

var billionaires = []billionaire{
	{Id: 1, Name: "Vladimir Putin", Country: "Russia ", NetWorth: 500.11, Image: "https://static1.therichestimages.com/wordpress/wp-content/uploads/2021/03/Vladimir-Putin-Net-Worth.jpg?q=50&fit=crop&w=60&h=60&dpr=1.5"},
	{Id: 2, Name: "Jeff Bezos", Country: "USA", NetWorth: 200.32, Image: "https://static1.therichestimages.com/wordpress/wp-content/uploads/Jeff-Bezos1.jpg?q=50&fit=crop&w=60&h=60&dpr=1.5"},
	{Id: 3, Name: "Elon Musk", Country: "USA", NetWorth: 300.45, Image: "https://static1.therichestimages.com/wordpress/wp-content/uploads/2012/03/Elon-Musk.jpg?q=50&fit=crop&w=60&h=60&dpr=1.5"},
	{Id: 4, Name: "The Mars Family", Country: "USA", NetWorth: 276.54, Image: "https://static1.therichestimages.com/wordpress/wp-content/uploads/2021/03/Mars-Family-Net-Worth.jpg?q=50&fit=crop&w=60&h=60&dpr=1.5"},
	{Id: 5, Name: "Bill Gates", Country: "USA", NetWorth: 298.65, Image: "https://static1.therichestimages.com/wordpress/wp-content/uploads/Bill-Gates2.jpg?q=50&fit=crop&w=60&h=60&dpr=1.5"},
}

func getBillionaires(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, billionaires)

}

// postBillionaire adds a billionaire from json received in the request body
func postBillionaires(context *gin.Context) {
	var newBillionaire billionaire
	// call BindJson to bind the received Json to newBillionaire
	if err := context.BindJSON(&newBillionaire); err != nil {
		return
	}
	// Add the new album to the slice
	billionaires = append(billionaires, newBillionaire)
	context.IndentedJSON(http.StatusCreated, newBillionaire)
}

func main() {
	router := gin.Default()
	router.GET("/billionaires", getBillionaires)
	router.POST("/billionaires", postBillionaires)
	router.Run("localhost:8080")
}
