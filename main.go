// pkgm + Tab + Tab

package main

import (
	"fmt"
	"math"
)

//Basic custom function, define type of input
func printInput(input string) {
	fmt.Println(input)
}

func printInput2(input string) string {
	return input
} //return type must be declared

//Return multible values
func returnMultibleValues(input1 int, input2 int) (multiply int, sum int) {
	multiply = input1 * input2
	sum = input1 + input2
	return
}

// struct shortcut => tys + Tab + Tab
type person struct {
	name string
	age  int
}
type developer struct {
	person     // developer struct inherited from person struct
	experience int
}

func (i person) printPerson() {
	fmt.Println(i)
}

//Interface
type shape interface {
	area() float64 //interface uses same named function as area()
}
type rectangular struct {
	longSide, shortSide float64
}
type circle struct {
	radius float64
}

// function for reached via interface
func (r rectangular) area() float64 {
	return r.longSide * r.shortSide
}

// function for reached via interface
func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
	//return 3.14 * c.radius * c.radius
}

// function for using interface as parameter
func shapeArea(i shape) {
	fmt.Println(i.area())
}

//struct for Maps
type City struct {
	plate, population int
}

func main() {
	// Doesn't need ";" end of line

	// fmt a library or package for using Println()
	fmt.Println("Hello Go") // shortcut => fp + Tab + Tab
	//print formated value
	fmt.Printf("Formatten string value %d", 5)
	fmt.Println("")

	//VARIABLES
	// variables don't need to declared type but if it needed it is allowed. For instance float32 or float64
	var num1 = 10 //didn't declared type as int
	fmt.Println(num1)
	var num2 float64 = 10.5 // declared type as float64
	fmt.Println(num2)
	var num3 float64
	fmt.Println(num3) // if variable didn't have a value it will return default value => 0
	num4 := 10        // easiest way for creating variable with :=
	fmt.Println(num4)
	var text = "Hello Go"
	fmt.Println(text)
	var bool1 = true
	fmt.Println(bool1)

	//CONSTANTS
	const constant1 = "Constant doesn't need a type"
	fmt.Println(constant1)
	const constant2 float64 = 10.5 // but if types needed they are allowed. For instance float32 or float64
	fmt.Println(constant2)

	//ARRAYS
	var array1 = [6]int{1, 2, 3, 4}      //[6] is length of array. 5th and 6th indexes given as default value
	var array2 = []string{"1", "2", "3"} //[] don't need to give length for array.
	fmt.Println(array1)
	fmt.Println(array2)

	//SLICES
	//Slices are parts of an array
	var slice1 = array1[0:3] // (0,3] => 1,2,3
	fmt.Println(slice1)
	var slice2 = array1[:3] // (0,3] => 1,2,3 same with [0:3]
	fmt.Println(slice2)
	var slice3 = array1[1:] // (1,end] => 2,3,4,0,0 starts with first index ends with last element of array
	fmt.Println(slice3)

	//MAKE FUNCTIONS
	//make() functions create slice,map or channel
	make1 := make([]int, 5) // make() function created an int array with 5 elements => 0,0,0,0,0
	fmt.Println(make1)
	make1[2] = 2
	fmt.Println(make1)

	//MAPS
	//Similar like multidimensional arrays
	//Maps can be created with Make() function
	cityMap := make(map[string]City) //Create map with key:string , value: City
	cityMap["Ankara"] = City{plate: 06, population: 12000000}
	cityMap["Istanbul"] = City{plate: 34, population: 18000000}
	cityMap["Izmir"] = City{plate: 35, population: 8000000}
	fmt.Println(cityMap)

	//FUNCTIONS
	printInput("printInput1")
	fmt.Println(printInput2("printInput2"))
	multiply, sum := returnMultibleValues(5, 10) //Returns two different values
	fmt.Println(multiply, sum)
	multiply, _ = returnMultibleValues(5, 10) //Returns one parameter _ is null variable
	fmt.Println(multiply)

	//INPUT FROM TERMINAL
	//bufio library allows to use input from terminal
	/*
		inputFromTerminal := bufio.NewReader(os.Stdin)//writes input in terminal
		inputFromUser, _ := inputFromTerminal.ReadString('\n') // reads input from terminal
		fmt.Println(inputFromUser)
	*/

	//IF-ELSE
	num5 := 10
	if num5 == 10 {
		fmt.Println(true)
	} else if num5 == 5 {
		fmt.Println(false)
	} else {
		fmt.Println("false")
	}

	//LOOPS
	//There are two types loops as for and for range loops
	//for loops
	count := 10
	for i := 0; i < count; i++ {
		fmt.Print(i)
	}
	fmt.Println("")
	//Another versioon of for loop default index starts with 0 as i:=0;i<count;i++
	i := 0
	for i < count {
		fmt.Print(i)
		i++
	}
	fmt.Println("")
	//Break usage in for loop
	for i := 0; i < count; i++ {
		if i == 5 {
			break
		} else {
			fmt.Print(i)
		}
	}
	fmt.Println("")
	// for range loop
	array3 := []string{"Java", "C#", "JavaScript", "CSS", "HTML", "GO"}
	for i, d := range array3 {
		fmt.Println(i, d)
	}

	//POINTERS
	//Pointer refers different variables to same references on heap
	// & => gets Reference Key of variable
	// * => get Value with Referance Key
	pointer1 := "Pointer1 Main Value"
	fmt.Println(pointer1)  // pointer1
	fmt.Println(&pointer1) // & => gets reference key of variable. pointer1 reference key.
	pointer2 := &pointer1  // Reference of pointer2 refers referance key of pointer1. So pointer2 has own ref key but value is pointer1 value
	fmt.Println(*pointer2) // * => gets value of pointed reference key
	fmt.Println(pointer2)  // Value of pointer2 which is same with ref key of pointer1
	fmt.Println(&pointer2) // Reference key of pointer2 which is different than pointer1

	//OOP in GoLang
	// struct is similar with class structure
	person1 := person{name: "Erhan", age: 33} // person instance created from person struct
	//person2 := person{"Erhan", 33}
	fmt.Println(person1.name, person1.age)
	//Inheritance
	// developer1 := developer{person: person{name: person1.name, age: person1.age},experience: 3}
	developer1 := developer{person{name: person1.name, age: person1.age}, 3}
	fmt.Println(developer1.name, developer1.age, developer1.experience)
	//Functions with struct
	person1.printPerson() // printPerson uses person type object. it is using like extension method
	//Interface
	//Need to have an instance from interface. It is possible in Golang.
	var instanceFromInterface shape  //1-Have an instance from interface
	shape1 := rectangular{5, 3}      //2-Have a variable from struct
	shape2 := circle{5}              //Have a variable from struct
	instanceFromInterface = &shape1  //3-point interface reference to variable from struct
	shapeArea(instanceFromInterface) //4- call function for interface
	instanceFromInterface = &shape2  // point interface reference to variable from struct
	shapeArea(instanceFromInterface) // call function for interface

	//Concurrency
	//Parallelism
	//goroutine
	//Channnel

}
