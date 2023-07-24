package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Database_Con, Dberr = sql.Open("mysql", "root:password@tcp(localhost:3306)/bankData")

var Find_customer = "select name, ssn, account_no, amount from customers where account_no = ?"
var AllCustomerQuery = "select name, ssn, account_no, amount from customers"
var AddNewCustomer = "insert into customers(account_no,name,ssn,amount) values(?,?,?,?)"
var DeleteCustomerQuery = "delete from customers where account_no = ?"
var TotalCusCount = "select count(account_no) from customers"
var UpdateCustomerQuery = "update customers set amount = ? where account_no = ?"
