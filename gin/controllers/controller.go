package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"gin/models"

	db "gin/database"

	"github.com/gin-gonic/gin"
)

var Desired_Customer models.Customer

var Find_customer = func(ctx *gin.Context) {
	RCustomer, err := strconv.Atoi(ctx.Param("id"))
	if err == nil {
		customerRow := db.Database_Con.QueryRow(db.Find_customer, RCustomer)
		error := customerRow.Scan(&Desired_Customer.CName, &Desired_Customer.CSSN, &Desired_Customer.CAccNo, &Desired_Customer.Amount)
		if error == nil {
			ctx.JSON(http.StatusOK, Desired_Customer)
		}
	}

}

var AllCustomers = func(ctx *gin.Context) {
	aBank := models.Bank{}
	rows, err := db.Database_Con.Query(db.AllCustomerQuery)

	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&Desired_Customer.CName, &Desired_Customer.CSSN, &Desired_Customer.CAccNo, &Desired_Customer.Amount)

		if err == nil {
			aBank.GenCus = append(aBank.GenCus, Desired_Customer)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"customers": aBank.GenCus,
	})
	defer rows.Close()

}

var AddNewCustomer = func(ctx *gin.Context) {
	var new_customer models.Customer

	rand.NewSource(time.Now().Unix())
	new_customer.CAccNo = rand.Intn(20)
	err := ctx.ShouldBindJSON(&new_customer)
	fmt.Println(new_customer)
	if err != nil {
		fmt.Println(err.Error())
	}
	stmt, err := db.Database_Con.Prepare(db.AddNewCustomer)
	if err != nil {
		fmt.Println(err)
	}
	_, error := stmt.Exec(new_customer.CAccNo, new_customer.CName, new_customer.CSSN, new_customer.Amount)

	if error != nil {
		fmt.Println("Error in adding customer", err)
	} else {
		ctx.JSON(http.StatusOK, "New recorded inserted")
	}

}

func DeleteCustoemr(ctx *gin.Context) {
	delAccNum, _ := strconv.Atoi(ctx.Param("id"))

	stmt, err := db.Database_Con.Prepare(db.DeleteCustomerQuery)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, error := stmt.Exec(delAccNum)
	if error != nil {
		fmt.Println(err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}

func TotalCustomers(ctx *gin.Context) {
	var count int
	Row := db.Database_Con.QueryRow(db.TotalCusCount)
	Row.Scan(&count)
	ctx.JSON(http.StatusOK, gin.H{
		"number of Customers": count,
	})
}

func WithDraw(ctx *gin.Context) {

	User_Amount, _ := strconv.Atoi(ctx.Param("amount"))

	if Desired_Customer.Amount > User_Amount {
		Desired_Customer.Amount = Desired_Customer.Amount - User_Amount

		stmt, err := db.Database_Con.Prepare(db.UpdateCustomerQuery)
		if err != nil {
			fmt.Println(err)
		}
		_, error := stmt.Exec(Desired_Customer.Amount, Desired_Customer.CAccNo)
		if error != nil {
			fmt.Println(error)
		}
		ctx.JSON(http.StatusOK, "amount updated successfully")
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Error":                   "Not enough balance to process the request",
			"Present Account Balance": Desired_Customer.Amount,
		})
	}

}

func Deposit(ctx *gin.Context) {
	User_Amount, _ := strconv.Atoi(ctx.Param("amount"))
	Desired_Customer.Amount = Desired_Customer.Amount + User_Amount

	stmt, err := db.Database_Con.Prepare(db.UpdateCustomerQuery)
	if err != nil {
		fmt.Println(err)
	}
	_, error := stmt.Exec(Desired_Customer.Amount, Desired_Customer.CAccNo)
	if error != nil {
		fmt.Println(error)
	}
	ctx.JSON(http.StatusOK, "amount updated successfully")
}
