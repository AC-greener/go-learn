package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	name    string
	age     int
	contact contactInfo
}

func main() {

	// alex := person{"Alex", 20}
	// alex := person{age: 20, name: "Alex"}
	// var alex person
	// alex.name = "Alex"
	// alex.age = 20
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		name: "Jim",
		age:  30,
		contact: contactInfo{
			email:   "",
			zipCode: 12345,
		}}
	// {

	// 	jimPointer := &jim
	// 	jimPointer.updateName("tom")
	// }
	jim.updateName("to11m")
	fmt.Printf("%+v", jim)
	// fmt.Printf(jim)

} 

func (ponterToPerson *person) updateName(newFirstName string) {
	ponterToPerson.name = newFirstName
	//和上面的结果相等 (*ponterToPerson).name = newFirstName
}
