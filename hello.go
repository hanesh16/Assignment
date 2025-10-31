package main

import "fmt"

func day1Hello() {
	fmt.Println("Hello, world!")
}

func day2() {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println("Hello:", i)
		}
	}
}

func day3structs() {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Hanesh", Age: 24}
	fmt.Println("Person:", p.Name, "Age:", p.Age)
}

func day3pointers() {
	x := 10
	y := &x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", y)
	*y = 20
	fmt.Println("New value of x:", x)
}

func day3array() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1[1:4]
	fmt.Println("Array 1:", arr1)
	fmt.Println("Array 2:", arr2)
}

func day3switch() {
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Friday":
		fmt.Println("End of the work week")
	default:
		fmt.Println("Midweek")
	}
}

func day2for() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func day2defer() {
	defer fmt.Println("World")
	for i := 0; i < 2; i++ {
		fmt.Println("Hello", i)
	}
}

func day4slices() {
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)
	fmt.Println("Slice:", slice)
}

func day4maps() {
	m := make(map[string]int)
	m["Alice"] = 30
	m["Bob"] = 25

	for k, v := range m {
		fmt.Println(k, "is", v, "years old")
	}
}

func main() {
	day1Hello()
	day2()
	day2for()
	day2defer()
	day3structs()
	day3pointers()
	day3array()
	day3switch()
	day4slices()
	day4maps()
}
