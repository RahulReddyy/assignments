package main

import (
	"fmt"
	db "gin/database"
	"gin/mappings"
)

func main() {

	if db.Dberr != nil {
		fmt.Println("Database connection failed")
		return
	}
	mappings.UrlMappings()
	mappings.Router.Run()
}
