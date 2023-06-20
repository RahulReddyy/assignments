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

type ecommmerce struct {
	inventory map[string]int //key = product, value = price
	cart      map[string]int //key = product, value = count of product
}

func (c ecommmerce) addItem(item string) {
	for invenValue := range c.inventory {
		if invenValue == item {
			c.cart[invenValue]++
			fmt.Println("Item added Successfully")
		}
	}
}

func (c ecommmerce) decreaseCount(item string) {
	for invenValue := range c.inventory {
		if invenValue == item && c.cart[invenValue] > 0 {
			c.cart[invenValue]--
		}
	}
}

func (c ecommmerce) listCart() {
	var bill int
	for item := range c.cart {
		for invItem := range c.inventory {
			if item == invItem {
				fmt.Printf("You have selected %d %s each of Price %d$", c.cart[item], item, c.inventory[invItem])
				fmt.Println()
				price := c.inventory[invItem]
				count := c.cart[item]
				bill += count * price
				fmt.Println("Total Bill:", bill, "$")
			}
		}

	}
}

func main() {
	newCart := ecommmerce{
		inventory: map[string]int{
			"apple":  2,
			"mango":  3,
			"orange": 4,
			"yogurt": 9,
			"phone":  6,
		}, cart: make(map[string]int),
	}
	var SearchItem, invenItem string
	var matchCheck bool
	var flag, countItem int

	for {
		if SearchItem == "" {
			fmt.Println("Enter the search item ")
			fmt.Scan(&SearchItem)
			for item := range newCart.inventory {
				pattern := ".\\s*" + item + ".*" + "|" + item + ".*"
				reobj, _ := regexp.Compile(pattern)
				matchCheck = reobj.MatchString(SearchItem)
				if matchCheck {
					invenItem = item
					//fmt.Println("inven ", invenItem)
					break
				}
			}
			if matchCheck {
				fmt.Println("if you want add item to cart enter 1(number), to exit enter zero(number)")
				fmt.Scan(&flag)

				if flag == 1 {
					newCart.addItem(invenItem)
				}
			} else {
				fmt.Println("Searched item is not present in the inventory")
				return
			}
		}
		for {
			fmt.Println("if you want to increase count enter 1(number) \n press 0(number) to decrease count \n to return to adding another item enter 3")
			fmt.Scan(&countItem)
			if countItem == 1 {
				newCart.addItem(invenItem)
			} else if countItem == 0 {
				if newCart.cart[invenItem] > 0 {
					newCart.decreaseCount(invenItem)
				}
			} else if countItem == 3 {
				break
			}
		}

		fmt.Println("To add another item press 1 or to list billing details press 5")
		fmt.Scan(&flag)
		if flag == 1 {
			SearchItem = ""
		} else if flag == 5 {
			newCart.listCart()
			return
		}

		if flag == 0 || countItem == 0 {
			return
		}
	}

}
