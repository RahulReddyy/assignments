package main

import (
	"fmt"
)

type bank struct {
	genCus []customer
}

func (bvar *bank) noOfCustomers() int {
	return len(bvar.genCus)
}

func (bvar *bank) printData() {
	fmt.Println((bvar.genCus))
	for _, cusdetail := range bvar.genCus {
		fmt.Println("Customer Name", cusdetail.cName)
		fmt.Println("Customer AccountNumber", cusdetail.cAccNo)
		fmt.Println("Customer SSN", cusdetail.cSSN)
		fmt.Println("Customer Amount", cusdetail.amount)
		fmt.Println()
	}
}
func (bvar *bank) addCustomer() {

	newCus := &customer{}
	fmt.Print("Enter customer name : ")
	fmt.Scan(&newCus.cName)
	fmt.Print("Enter customer SSN : ")
	fmt.Scan(&newCus.cSSN)
	fmt.Print("Enter customer deposit amount : ")
	fmt.Scan(&newCus.amount)
	newAccNoAdd := bvar.noOfCustomers()
	newCus.cAccNo = newAccNoAdd + 1
	fmt.Println("Account number for the customer is : ", newCus.cAccNo)
	bvar.genCus = append(bvar.genCus, *newCus)
	fmt.Println(len(bvar.genCus))

}

func (bvar *bank) deleteCustomer(delAccNum int) {

	for index, delCus := range bvar.genCus {
		if delCus.cAccNo == delAccNum {
			bvar.genCus = append(bvar.genCus[:index], bvar.genCus[index:]...)
		}
	}
}

type customer struct {
	cName  string
	cAccNo int
	cSSN   int
	amount int
}

func (cus *customer) withDraw(withDrawAmount int) {
	if cus.amount >= withDrawAmount {
		cus.amount = cus.amount - withDrawAmount
		fmt.Println("Remaning amount", cus.amount)
	} else {
		fmt.Println("Not enough balance to withdraw")
	}
	fmt.Println("in withdraw :", cus)
}
func (cus *customer) depositMoney(dMoney int) {
	cus.amount += dMoney
	fmt.Println("Total balance in the account", cus.amount)
}
func custFind(tBank *bank, accNum int) *customer {
	var selectedCustomer customer
	for index, customerdata := range tBank.genCus {
		if customerdata.cAccNo == accNum {
			return &tBank.genCus[index]
		}
	}
	return &selectedCustomer
}
func main() {
	aBank := &bank{}
	customer1 := &customer{cName: "customer1", cAccNo: 0001, cSSN: 0011, amount: 1000}
	customer2 := &customer{cName: "customer2", cAccNo: 0002, cSSN: 0012, amount: 1001}
	customer3 := &customer{cName: "customer3", cAccNo: 0003, cSSN: 0012, amount: 1001}
	aBank.genCus = append(aBank.genCus, *customer1, *customer2, *customer3)

	var cusCheck, serviceType, accNum, amount, exitCheck int
	fmt.Print("if New customer press 1(number) else enter 0(number)")
	fmt.Scan(&cusCheck)
	for {
		if cusCheck == 1 {
			fmt.Println("hey")
			aBank.addCustomer()

		} else if cusCheck == 0 {
			fmt.Println("Enter the service number you require\n1.WithdrawMoney\n2.DepositMoney\n3.NumberOfCustomers\n4.PrintAllCustomers\n5.PrintBalance\n6.DeleteCustomer ")
			fmt.Scan(&serviceType)

			switch serviceType {
			case 1:
				fmt.Println("Enter account number")
				fmt.Scan(&accNum)
				fmt.Println("Enter amount to withDraw")
				fmt.Scan(&amount)
				(*custFind(aBank, accNum)).withDraw(amount)
				aBank.printData()
				amount = 0
			case 2:
				fmt.Println("Enter account number")
				fmt.Scan(&accNum)
				fmt.Println("Enter amount to Deposit")
				fmt.Scan(&amount)
				(*custFind(aBank, accNum)).depositMoney(amount)
				amount = 0
			case 3:
				fmt.Println("Number of customers in bank are", aBank.noOfCustomers())
			case 4:
				aBank.printData()
			case 5:
				fmt.Println("Enter account number")
				fmt.Scan(&accNum)
				fmt.Println("balance in the account is", (*custFind(aBank, accNum)).amount)
			case 6:
				fmt.Println("Enter account number")
				fmt.Scan(&accNum)
				aBank.deleteCustomer(accNum)
				fmt.Println("Account deleted successfully")
			default:
				fmt.Println("Enter number mentioned above")
			}

		}
		fmt.Println("If you want to exit, enter 1\n to add another customer enter 0\n to view cusomter services enter 3")
		fmt.Scan(&exitCheck)
		if exitCheck == 0 {
			cusCheck = 1
		} else if exitCheck == 3 {
			cusCheck = 0
		}
		if exitCheck == 1 {
			return
		}
	}

}
