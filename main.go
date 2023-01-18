package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type billionaire struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"imageUrl"`
	//Country  string
	//netWorth float64
}

var billionaires = []billionaire{
	{Id: 1, Name: "Vladimir Putin", Image: "https://static1.therichestimages.com/wordpress/wp-content/uploads/2021/03/Vladimir-Putin-Net-Worth.jpg?q=50&fit=crop&w=60&h=60&dpr=1.5"},
	{Id: 2, Name: "Jeff Bezos", Image: "https://static1.therichestimages.com/wordpress/wp-content/uploads/Jeff-Bezos1.jpg?q=50&fit=crop&w=60&h=60&dpr=1.5"},
	{Id: 3, Name: "Elon Musk", Image: "https://static1.therichestimages.com/wordpress/wp-content/uploads/2012/03/Elon-Musk.jpg?q=50&fit=crop&w=60&h=60&dpr=1.5"},
	{Id: 4, Name: "The Mars Family", Image: "https://static1.therichestimages.com/wordpress/wp-content/uploads/2021/03/Mars-Family-Net-Worth.jpg?q=50&fit=crop&w=60&h=60&dpr=1.5"},
	{Id: 5, Name: "Bill Gates", Image: "https://static1.therichestimages.com/wordpress/wp-content/uploads/Bill-Gates2.jpg?q=50&fit=crop&w=60&h=60&dpr=1.5"},
}

func getBillionaires(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, billionaires)

}

func main() {
	router := gin.Default()
	router.GET("/billionaires", getBillionaires)
	router.Run("localhost:9090")
}
