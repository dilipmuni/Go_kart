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

func asuser() {
	for !add_user() {

	}
	fmt.Println("1.View products\n2.Place Order\n3.View Cart\n4.Change Order\n5.Cancel Order\n 0.Logout")
	var choose int
	choose = 1
	for choose != 0 {
		fmt.Scanln(&choose)
		switch choose {
		case 0:
			return
		case 1:
			list_of_products()
		case 2:
			add_product()
		case 3:
			getorders()
		case 4:
			change_product()
		case 5:
			cancel_order()
		default:
			fmt.Println("Wrong Entry\n")
		}
	}
}

func asdeliveryrep() {

}

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
	fmt.Printf("\n")
	for _, v := range all_Products {
		fmt.Printf(" %s\n", v)
	}

}

type orders struct {
	name string
	qty  int
}

var my_orders []orders

func cancel_order() {
	my_orders = []orders{}
}

func add_product() bool {
	var a string
	fmt.Scanln(&a)

	fmt.Printf("Quantity : ")
	var b int
	fmt.Scanln(&b)

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

	if b == 0 {
		return false
	}
	ChangeOrders(a, b)
	fmt.Println(my_orders[:])
	return true
}

func ChangeOrders(name string, qty int) {
	for _, v := range my_orders {
		if v.name == name {
			v.qty = qty
		}
	}

}

func getorders() {
	fmt.Println(my_orders[:])
}

func main() {

	var loginas int

	for true {
		fmt.Printf("Login As\n1.User\n2.Delivery Agent\n")
		fmt.Scanln(&loginas)
		switch loginas {
		case 1:
			asuser()
		case 2:
			asdeliveryrep()
		default:
			fmt.Println("Please login with proper credentials")
		}
	}

}
