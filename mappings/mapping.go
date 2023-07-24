package mappings

import (
	"ginapi/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func UrlMappings() {
	Router = gin.Default()

	v1 := Router.Group("")
	{
		v1.GET("/findcustomer/:id", controllers.Find_customer)
		v1.GET("/allcustomers", controllers.AllCustomers)
		v1.POST("/add_newcustomer", controllers.AddNewCustomer)
		v1.DELETE("/deletecustomer/:id", controllers.DeleteCustoemr)
		v1.GET("/totalcustomers", controllers.TotalCustomers)
		v1.PATCH("/withdraw/:amount", controllers.WithDraw)
		v1.PATCH("/deposit/:amount", controllers.Deposit)

	}
}
