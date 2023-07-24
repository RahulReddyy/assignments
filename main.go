package main

import (
	"fmt"
	db "ginapi/database"
	"ginapi/mappings"
)

func main() {

	if db.Dberr != nil {
		fmt.Println("Database connection failed")
		return
	}
	fmt.Println(db.Dberr)
	mappings.UrlMappings()

	mappings.Router.Run(":8080")
}
