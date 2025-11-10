package main

import (
	"fmt"
	"sync"
	"time"
)

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

func day5palindrome() {
	word := "madam"
	isPalindrome := true
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println(word, "is a palindrome")
	} else {
		fmt.Println(word, "is not a palindrome")
	}
}

func day5fibonacci() {
	n := 10
	a, b := 0, 1
	fmt.Print("Fibonacci series: ")
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

func day5maxnumarray() {
	arr := []int{3, 5, 2, 8, 1}
	max := arr[0]
	for i, z := range arr {
		if z > max {
			max = z
		}
		fmt.Println("Index:", i, "Value:", z)
	}
	fmt.Println("Maximum number in array is:", max)
}

func day6goroutine() {
	go func() {
		fmt.Println("Hello from goroutine!")
	}()

	fmt.Println("Hello from main!")

	time.Sleep(100 * time.Millisecond)
}

func day6waitgroup() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Number:", i)
		}
		wg.Done()
	}()

	go func() {
		letters := []string{"A", "B", "C", "D", "E"}
		for _, letter := range letters {
			fmt.Println("Letter:", letter)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Both goroutines finished")
}

func day6channel() {

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Number from goroutine:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(600 * time.Millisecond)
	fmt.Println("Done printing numbers")
}

func day6timeout() {

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Task completed after 3 seconds")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Main gave up waiting (timeout)")
}

func day6worker() {
	var result int
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		number := 5
		result = number * 2
		fmt.Println("Worker calculated:", result)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final result:", result)
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
	day5palindrome()
	day5fibonacci()
	day5maxnumarray()
	day6goroutine()
	day6waitgroup()
	day6channel()
	day6timeout()
	day6worker()
}
