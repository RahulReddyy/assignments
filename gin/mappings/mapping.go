package mappings

import (
	"gin/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func UrlMappings() {
	Router = gin.Default()

	v1 := Router.Group("")
	{
		v1.GET("/find_customer/:id", controllers.Find_customer)
		v1.GET("/all_customers", controllers.AllCustomers)
		v1.POST("/add_new_customer", controllers.AddNewCustomer)
		v1.DELETE("/delete_customer/:id", controllers.DeleteCustoemr)
		v1.GET("/total_customers", controllers.TotalCustomers)
		v1.PATCH("/withdraw/:amount", controllers.WithDraw)
		v1.PATCH("/deposit/:amount", controllers.Deposit)

	}
}
