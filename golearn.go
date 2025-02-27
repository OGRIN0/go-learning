package main

import (
	"fmt"
)

// 1st program
// func main(){
// 	fmt.Println("Hello, World")
// }




// 2nd program
// func main(){
// 	var nameOne string = "vanakam"
// 	var nameTwo = "helllo"
// 	var nameThree string

// 	fmt.Println(nameOne, nameTwo, nameThree)

// 	nameOne = "Go"
// 	nameThree = "Golang"

// 	fmt.Println(nameOne, nameTwo, nameThree)
// 	nameFour := "goshi"
// 	fmt.Println(nameFour)

// 	var ageOne int = 10
// 	var ageTwo = 30
// 	var ageThree = 50

// 	fmt.Println(ageOne, ageTwo, ageThree)

// 	var numOne int8 = 110
// 	var numTwo int64 = 7346287676

// 	fmt.Println(numOne, numTwo)
// }




// 3rd program
// func main(){
// 	var name string = "Luke"
// 	var age int = 30
// 	fmt.Print("Hello")
// 	fmt.Print("\n")

// 	fmt.Println("Hello", name)
// 	fmt.Printf("Hi I'm %v and I'm %v years old! \n", name, age)

// 	var stringOne = fmt.Sprintf("Hi I'm %v and I'm %v years old! \n", name, age)
// 	fmt.Println(stringOne)
// }

// 4th program
// func main() {
// 	var ages [3]int = [3]int{20, 30, 50}

// 	names := [4]string{"luke", "bake", "coke", "dake"}

// 	fmt.Println(ages, len(ages))
// 	fmt.Println(names[0:3], len(names))

// 	var scores = []int{100, 50, 60}
// 	scores[2] = 50
// 	scores = append(scores, 30)

// 	fmt.Println(scores, len(scores))
// }




// 5th program
// func main(){
// 	greeting := "Hello Everyone!" 

// 	fmt.Println(strings.Contains(greeting, "Hi!"))
// 	Hi := fmt.Sprintln(strings.ReplaceAll(greeting, "Hello", "Hi!"))
// 	fmt.Println(strings.ToUpper(greeting))

// 	fmt.Println("original string values =", greeting)
// 	fmt.Println("modified greeting =", Hi)

// 	// fmt.Println(string.Split(greeting, " "))

// 	ages := []int{56, 64, 21, 54, 80, 32}
// 	sort.Ints(ages)
// 	fmt.Println(ages)


// 	index := sort.SearchInts(ages, 56)
// 	fmt.Println(index)
// }




// 6th program
// func main(){
// 	x := 1
// 	for x < 4 {
// 		fmt.Println("value of x is", x)
// 		x++
// 	}

// 	for i := 1; i<3; i++{
// 		fmt.Println("value of i is", i)
// 	}

// 	names := []string{"gokey", "Hiren", "Vegeta", "Dragon"}

// 	for i := 0; i < len(names); i++{
// 		fmt.Println(names[i])
// 	}

// 	for index, value := range names{
// 		fmt.Printf("the value at index %v is %v \n", index, value)
// 	}
// }




func main(){
	age := 30 

	fmt.Println(age==30)
	fmt.Println(age!=30)

	if age < 20{
		fmt.Println("Not eligible to stand")
	} else {
		fmt.Println("Eligible to vote")
	}

	names := []string{"gokey", "Hire", "Vegeta", "Roady"}

	for index, value := range names{
		if index == 0{
			fmt.Println("Continuing with you Goku your pos is", index)
			continue
		}
		if index ==3{
			fmt.Println("Sorry Roady, I'm breaking with Vegeta your number is not available", index)
			break
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
}



