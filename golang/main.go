package main

import "fmt"
import "math"

import (
"strings"
"sort"
"os"
"log"
"io/ioutil"
"time"
"net/http"
)

func main() {
	fmt.Println("Hello World!")

	var age int = 40

	var score float32 = 93.27

	birthYear := 2012

	fmt.Println(age, score, birthYear)

	fmt.Println(age+birthYear)

	const pi float32 = 3.14159

	var (
		varA = 2
		varB = 3
		)

	fmt.Println(varA+varB)

	var firstName string = "Abhiman Sitaram"
	var lastName string = " Kolte"

	fmt.Println(len(firstName))

	fmt.Println(firstName + lastName)

	fmt.Printf("%d %f %d %s %s", age, score, birthYear, firstName, lastName)

	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j, "Children")
	}

	if age > 18 {
		fmt.Println("Can Drive")
	} else {
		fmt.Println("Can't Drive")
	}

	switch age {
		case 16: fmt.Println("Drive")
		case 18: fmt.Println("Vote")
		default: fmt.Println("Fun")
	}

	var arrayInts[10] int

	for k := 0; k < 10; k++ {
		arrayInts[k] = k
	}

	favNums := [2]int32 {1, 2}

	for l, value := range arrayInts {
		fmt.Println(value, l)
	}

	for _, value := range favNums {
		fmt.Println(value)
	}

	numSlice := []int {1,2,3,4,5}


	numSlice2 := numSlice[1:3]

	fmt.Println(numSlice)

	fmt.Println(numSlice2)

	friendAge := make(map[string] int)

	friendAge["Abhiman"] = 22
	friendAge["Rajas"] = 25
	friendAge["Saha"] = 29

	fmt.Println(friendAge)
	fmt.Println(len(friendAge))


	listNums := []int {1,2,3,4,5,6,7,8,9,10}

	fmt.Println(add(listNums))

	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)


	fmt.Println(sub(1,2,3,4,5,6,7,8,9))

	num3 := 3


	square := func() int {
		num3 *= num3

		return num3
	}

	fmt.Println(square())
	fmt.Println(square())
	fmt.Println(square())


	fmt.Println(fact(1))
	fmt.Println(fact(2))
	fmt.Println(fact(3))
	fmt.Println(fact(0))
	fmt.Println(fact(20))


	fmt.Println(safeDivision(1,0))

	defer printTwo()
	printOne()


	demPanic()


	x := 2
	fmt.Println(x)

	changeVal(&x)

	fmt.Println(x)

	rect1 := Rectangle{0,50,10,10}

	fmt.Println(rect1)

	fmt.Println(rect1.area())


	rect := Rectangle{1,1,20,50}
	circle := Circle{1}

	fmt.Println(getArea(rect))
	fmt.Println(getArea(circle))

	fmt.Println(strings.Contains("Abhiman Kolte", "Abhi"))
	fmt.Println(strings.Index("Abhiman Kolte", "Abhima"))
	fmt.Println(strings.Replace("abhiman Kolte", "a", "b", 100))

	csvString := "1,2,3,4,5,6,7,8,9,10"

	fmt.Println(strings.Split(csvString, ","))

	listOfLetters := []string {"a", "q", "d","e"}

	sort.Strings(listOfLetters)

	fmt.Println(listOfLetters)


	file, err := os.Create("a.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Lol this is Golang")

	file.Close()

	stream, err := ioutil.ReadFile("a.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(stream))


	// Web Server Basics

	//http.HandleFunc("/", handler)
	//http.HandleFunc("/earth", handler2)

	//http.ListenAndServe(":8080", nil)


	for z := 0; i < 10; z++ {
		go count(z)
	}

	time.Sleep(time.Millisecond * 11000)


}


func count(id int) {
	for i := 0; i < 10; i++ {
			fmt.Println(id, ":", i)
		time.Sleep(time.Millisecond * 1000)
	}
}


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Earth Page!\n")
}

func sub(args ...int) int {
	result := 0

	for _, value := range args {
		result -= value
	}
	return result
}


func next2Values(number int) (int, int) {
	return number+1, number+2
}


func add(nums []int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

func fact(num int) int {
	if num == 0 {
		return 1
	}

	return num * fact(num - 1)
}


func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }


func safeDivision(num1, num2 int) int {

	defer func() {
		fmt.Println(recover())
	}()

	solution := num1/num2
	return solution
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
		}()
		panic("PANIC")
}


func changeVal(x *int) {
	*x = 100
}


type Rectangle struct {
	leftX float64
	topY float64
	height float64
	width float64
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}


type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

