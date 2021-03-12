package main

import "fmt"

type User struct {
	name  string
	num   int
	place string
}

type Product struct {
	name string
}

var slc []User
var num_user int

func add_user() bool {

	fmt.Printf("Please enter your credentials\n")
	var u User
	fmt.Printf("Enter your name : ")
	fmt.Scanln(&u.name)

	fmt.Printf("Enter your Mobile Number : ")
	fmt.Scanln(&u.num)

	fmt.Printf("Enter your Place : ")
	fmt.Scanln(&u.place)

	return AddUsers(u)
}

func AddUsers(user User) bool {

	if user.num < 6999999999 || user.num > 9999999999 {
		fmt.Println("Error : Mobile Number should be 10 digits")
		return false
	}
	for k, v := range slc {
		fmt.Println(k)
		if v.num == user.num {
			fmt.Println("Mobile number already Exist")
			return false
		}
	}
	slc = append(slc, user)

	return true
}

var all_Products [3]string = [3]string{"TV", "fridge", "Mobile"}

func list_of_products() {
	for k, v := range all_Products {
		fmt.Printf("%d . %s\n", k, v)
	}
	fmt.Printf("Enter 0 at anytime to come out\n")
}

type orders struct {
	name string
	qty  int
}

var my_orders []orders

func add_product() bool {

	var a string
	fmt.Scanln(&a)

	fmt.Printf("Quantity : ")
	var b int
	fmt.Scanln(&b)

	if a == "0" {
		return false
	}
	if b == 0 {
		return false
	}
	PlaceOrder(a, b)
	fmt.Println(my_orders[:])
	return true
}

func PlaceOrder(name string, qty int) {
	var u orders
	u.name = name
	u.qty = qty
	my_orders = append(my_orders, u)
}

func change_product() bool {
	fmt.Printf("Enter the name of the Product you want to change : ")
	var a string
	fmt.Scanln(&a)

	fmt.Printf("Quantity : ")
	var b int
	fmt.Scanln(&b)
	if a == "0" {
		return false
	}
	if b == 0 {
		return false
	}
	ChangeOrders(a, b)
	fmt.Println(my_orders[:])
	return true
}

func ChangeOrders(name string, qty int) {
	for k, v := range my_orders {
		if v.name == name {
			v.qty = qty
			fmt.Println(k)
		}
	}
}

func getorders() {
	fmt.Println(my_orders[:])
}

func main() {

	fmt.Printf("please Login\n")
	for !add_user() {

	}
	list_of_products()
	fmt.Printf("please enter your product : ")
	for add_product() {
	}
	for change_product() {

	}
	fmt.Println(my_orders[:])
}
