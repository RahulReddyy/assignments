package main

import (
	"fmt"
	"regexp"
)

// E-commerce
// 	inventory
// 	cart functionality-> adding,deleting, select more count, decrease count
// 	buy item
// 	list all items
// 	display total bill

type details struct {
	Price    int
	Quantity int
}

type ecommmerce struct {
	inventory map[string]details //item : {price,quantity}
	cart      map[string]int  // item : count
}

func (c ecommmerce) addItem(item string) {
	for invenValue := range c.inventory {
		if invenValue == item && c.inventory[invenValue].Quantity > 0 {
			quantValue := c.inventory[invenValue]
			quantValue.Quantity -= 1
			c.inventory[invenValue] = quantValue
			c.cart[invenValue]++
			fmt.Println("Item added Successfully")
		}
	}
}

func (c ecommmerce) decreaseCount(item string) {

	if c.cart[item] > 1 {
		for invenValue := range c.inventory {
			if invenValue == item {
				quantValue := c.inventory[invenValue]
				quantValue.Quantity += 1
				c.inventory[invenValue] = quantValue
				c.cart[invenValue]--
			}
		}
	} else if c.cart[item] == 1 {
		delete(c.cart, item)
		fmt.Println("Item deleted completely")
	}
}

func (c ecommmerce) listCart() {
	var countcheck int
	var itemCheck string
	var editCheck string
	for item := range c.cart {
		for invItem := range c.inventory {
			if item == invItem {
				fmt.Printf("You have selected %d %s each of Price %d$", c.cart[item], item, c.inventory[invItem].Price)
			}
		}

	}
	fmt.Println(".....Do you want edit cart, enter yes or no....")
	fmt.Scan(&editCheck)
	if editCheck == "yes" {
		fmt.Println("enter the item you want to increase or decrease count")
		fmt.Scan(&itemCheck)
		for invItem := range c.inventory {
			if itemCheck == invItem {
				fmt.Println("if you want to increase count enter 1 else enter to decrease count 3 ")
				fmt.Scan(&countcheck)
				if countcheck == 1 {
					c.addItem(invItem)
				} else if countcheck == 3 {
					if c.cart[invItem] > 0 {
						c.decreaseCount(invItem)
					} else {
						fmt.Println("Item not found")
					}
				}
			}
		}

	}

}

func (c ecommmerce) cartBill() {
	var bill int
	if len(c.cart) > 0 {
		fmt.Println(".....Billing Details....")
		for item := range c.cart {
			for invItem := range c.inventory {
				if item == invItem {
					fmt.Printf("You have selected %d %s each of Price %d$", c.cart[item], item, c.inventory[invItem].Price)
					fmt.Println()
					price := c.inventory[invItem].Price
					count := c.cart[item]
					bill += count * price
				}
			}

		}
		fmt.Println("Total Bill:", bill, "$")
		fmt.Println()
	} else {
		fmt.Println("Cart is empty !!!")
	}

}

func main() {
	newCart := ecommmerce{
		inventory: map[string]details{
			"apple": {
				Price:    4,
				Quantity: 7,
			},
			"mango": {
				Price:    5,
				Quantity: 7,
			},
			"orange": {
				Price:    6,
				Quantity: 7,
			},
			"yogurt": {
				Price:    7,
				Quantity: 7,
			},
			"phone": {
				Price:    8,
				Quantity: 7,
			},
		}, cart: make(map[string]int), 
	}
	var (
		SearchItem, invenItem string
		matchCheck, nextItem1 bool
		flag, countItem       int = 8, 9
	)

	for {
		if SearchItem == "" {
			fmt.Println("Enter the search item ")
			fmt.Scan(&SearchItem)
			for item := range newCart.inventory {
				pattern := "[a-zA-Z]]\\s*" + item + ".*" + "|" + item + ".*"
				reobj, _ := regexp.Compile(pattern)
				matchCheck = reobj.MatchString(SearchItem)
				if matchCheck {
					invenItem = item
					break
				}
			}
			if matchCheck {
				fmt.Printf("Item %s Found, Do you want add item to cart? \nIf yes, enter 1(number)", invenItem)
				fmt.Scan(&flag)
				if flag == 1 {
					newCart.addItem(invenItem)
					nextItem1 = true
				}
				for nextItem1 {
					fmt.Println()
					fmt.Print("To increase count enter 1(number)\nTo decrease count enter 5(number)\n")
					fmt.Print("To add another item 3(number)\nTo list billing details enter 7(number)\nTo list cart enter 11")
					fmt.Println()
					fmt.Println()
					fmt.Scan(&countItem)

					if countItem == 1 {
						newCart.addItem(invenItem)
					} else if countItem == 5 {
						if newCart.cart[invenItem] > 0 {
							newCart.decreaseCount(invenItem)
							fmt.Println()
						} else {
							fmt.Println("Item not found, Please add another Item")
						}

					} else if countItem == 3 {
						SearchItem = ""
						break
					} else if countItem == 11 {
						newCart.listCart()
						fmt.Println("To add another item 3(number)\nTo exit enter 12")
						fmt.Scan(&countItem)
						if countItem == 3 {
							SearchItem = ""
							break
						} else if countItem == 12 {
							return
						}
					} else if countItem == 7 {
						newCart.cartBill()
					}
				}

			}

		}

		if !matchCheck {
			fmt.Println("Searched item is not present in the inventory, Please search again")
			SearchItem = ""
		}
	}
}
