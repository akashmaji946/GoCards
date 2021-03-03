package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

type contactInfo struct {
	email   string
	phone   string
	pinCode int
}

type employee struct {
	age     int
	company string
	details person
	contact contactInfo
}

type doctor struct {
	age        int
	name       string
	specialist bool
	salary     float64
}

func (d doctor) print() {
	fmt.Printf("%+v\n", d)

}

func (d doctor) updateAge(newAge int) {
	d.age = newAge

}

func (d *doctor) updateName(newName string) {
	(*d).name = newName ////this is the right way, however we can omit '*'
}

func main() {
	fmt.Println("Hello from structs file")

	// // order of args must match sequence of params in defn
	// akash := person{"Akash", "Maji"}
	// fmt.Println(akash)

	// suraj := person{"Suraj", "Maji"}
	// fmt.Println(suraj)

	// // sequenced parameters: ordered
	// anup := person{firstName: "Anup", lastName: "Maji"}
	// fmt.Println(anup)

	// // named parameter: any order
	// mithu := person{lastName: "Maji", firstName: "Mithu"}
	// fmt.Println(mithu)

	// // another way of defining structs
	// var akash person
	// fmt.Println(akash)
	// fmt.Printf("%+v\n", akash)

	// akash.firstName = "Akash"
	// akash.lastName = "Maji"

	// fmt.Println(akash)
	// fmt.Printf("%+v\n", akash)

	// emp := employee{
	// 	age:     25,
	// 	company: "TCS",
	// 	details: person{
	// 		firstName: "Raju",
	// 		lastName:  "Sri",
	// 	},
	// 	contact: contactInfo{
	// 		email:   "raju@tcs.com",
	// 		phone:   "12345678",
	// 		pinCode: 12332,
	// 	},
	// }

	// fmt.Printf("%+v\n", emp)

	// munnabhai := doctor{
	// 	age:        32,
	// 	name:       "Sanjay Dutt",
	// 	specialist: true,
	// 	salary:     200000,
	// }

	// munnabhai.print()
	// munnabhai.updateAge(33) //not updated, use pointers (not local copy)
	// munnabhai.print()

	// munnabhai.updateName("Sanju Dutta")
	// (&munnabhai).print() //this is the right way, however we can omit '&'

	greets := []string{"hello", "hola", "thank you", "welcome"}
	fmt.Println(greets)
	greets[0] = "Hello, world!"
	fmt.Println(greets)

	update(greets, 1, "hola hola ho")
	fmt.Println(greets)

}

func update(ss []string, i int, s string) {
	ss[i] = s //this will update?? yes, but why? we dont have passed by pointers
}
