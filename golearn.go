package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"
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

// 7th program
// func main(){
// 	age := 30

// 	fmt.Println(age==30)
// 	fmt.Println(age!=30)

// 	if age < 20{
// 		fmt.Println("Not eligible to stand")
// 	} else {
// 		fmt.Println("Eligible to vote")
// 	}

// 	names := []string{"gokey", "Hire", "Vegeta", "Roady"}

// 	for index, value := range names{
// 		if index == 0{
// 			fmt.Println("Continuing with you Goku your pos is", index)
// 			continue
// 		}
// 		if index ==3{
// 			fmt.Println("Sorry Roady, I'm breaking with Vegeta your number is not available", index)
// 			break
// 		}

// 		fmt.Printf("the value at pos %v is %v \n", index, value)
// 	}
// }




// 8th program
// func sayGreeting(n string){
// 	fmt.Printf("Good Mornign %v \n", n)
// }

// func sayBye(n string){
// 	fmt.Printf("Good Bye %v \n", n)
// }

// func cycleNames(n []string, f func(string)){
// 	for _, v := range n{
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64{
// 	return math.Pi * r * r 
// }

// func main(){
// 	sayGreeting("giren")
// 	sayGreeting("koko")
// 	sayBye("Thanos")

// 	cycleNames([]string{"giren", "koko", "Thanos"}, sayGreeting)
// 	cycleNames([]string{"giren", "koko", "Thanos"}, sayBye)

// 	area1 := circleArea(5.5)
// 	area2 := circleArea(5.8)
// 	fmt.Println(area1, area2)
// 	fmt.Printf("circle 1 area %0.4f and cirecle 2 area is %0.4f \n", area1, area2)
// }




// 9th program
// func getInitials(n string) (string, string) {
// 	s := strings.ToUpper(n)
// 	names := strings.Split(s, " ")

// 	var initials []string 
// 	for _, v := range names {
// 		initials = append(initials, v[:1])
// 	}

// 	if len(initials)>1{
// 		return initials[0], initials[1]
// 	}

// 	return initials[0], "_"
// }




// 10th program
// func main(){
// 	fn, sn := getInitials("Thanos is the best living being_He thinks for everyone!")
// 	fmt.Println(fn, sn)

// 	an1, dn1 := getInitials("Giren Savog")
// 	fmt.Println(an1, dn1)

// }




// 11th program
// func main(){
// 	localScore := 34.7532  

// 	sayHi("Giren")

// 	for _, v := range points {  
// 		fmt.Println(v)
// 	}

// 	showScore(localScore)
// }




// 12th program
// func main(){
// 	menu := map[string]float64{
// 		"soup" :     4.99,
// 		"pie" :      5.65,
// 		"salad" :    6.43,
// 		"toffee" :   5.32,
// 	}

// 	fmt.Println(menu)
// 	fmt.Println(menu["pie"])

// 	for k, v := range menu{
// 		fmt.Println(k, "-", v)
// 	}
	
// 	phonebook := map[int]string{
// 		37654376257: "voki",
// 		84678526263: "goku",
// 		24165436272: "akon",
// 	}

// 	fmt.Println(phonebook)
// 	fmt.Println(phonebook[37654376257])

// 	phonebook[37654376258] = "giren"
// 	fmt.Println(phonebook)
// }




// 13th program
// func updateName(x string) string{
// 	x = "wedge"
// 	return x 
// }

// func updateMenu(y map[string]float64) {
// 	y["coffee"] = 2.88
// }

// func main() {
// 	name := "elon"

// 	updateName(name)

// 	fmt.Println(name)

// 	menu := map[string]float64{
// 		"apple pie" :   4.53245,
// 		"unicorn salad" : 234.342333,
// 	}

// 	updateMenu(menu)
// 	fmt.Println(menu)
// }





// 14th program
// func updateName(x *string){
// 	*x = "wedge"
// }

// func main(){
// 	name := "musk"

// 	// updateName(name)

// 	fmt.Println("memory address of the musk is", &name)

// 	m := &name 
// 	a := *m
// 	// fmt.Println(name)
// 	// fmt.Println(*m)

// 	updateName(a)
// 	fmt.Println(name)
// }




// 15th programs
// func main(){
// 	mybill := newBill("giren's bill")


// 	mybill.addItem("coke", 50.43)
// 	mybill.addItem("cokecola", 2.34)
// 	mybill.addItem("flowercake", 3.43)
// 	mybill.updateTip(2.32)
	

// 	fmt.Println(mybill.format())
// }



func getInput(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}


func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name:")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create a new bill name:", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOption(b bill) {
    reader := bufio.NewReader(os.Stdin)

    opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
    
    switch opt {
    case "a":
        name, _ := getInput("Item name: ", reader)
        price, _ := getInput("Item price: ", reader)
        
        p, err := strconv.ParseFloat(price, 64)
        if err != nil {
            fmt.Println("The price must be a number")
            promptOption(b)
            return
        }
        
        b.addItem(name, p)
        fmt.Println("Item added -", name, p)
        promptOption(b)
    case "t":
        tip, _ := getInput("Enter tip amount: ", reader)
        t, err := strconv.ParseFloat(tip, 64)
        if err != nil {
            fmt.Println("The tip must be a number")
            promptOption(b)
            return
        }
        b.updateTip(t)
        fmt.Println("Tip added:", t)
        promptOption(b)
    case "s":
        b.save()  // Call the save method
        fmt.Println(b.format())
        return
    default:
        fmt.Println("That was not a valid option...")
        promptOption(b)
    }
}

func main(){
	mybill := createBill()
	promptOption(mybill)

	fmt.Println(mybill)
}
