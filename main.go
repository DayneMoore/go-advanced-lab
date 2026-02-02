package main

import (
	"fmt"
	"os"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("Error: factorial is not defined for negative numbers")
	} else if n == 0 {
		return 1, nil
	}
	result, err := Factorial(n - 1)
	if err != nil {
		return 0, err
	}
	return n * result, nil
}

func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, fmt.Errorf("Error: prime check requires number >= 2")
	} else if n == 2 {
		return true, nil
	} else if n%2 == 0 {
		return false, nil
	} else {
		for i := 3; i*i <= n; i += 2 {
			if n%i == 0 {
				return false, nil
			}
		}
		return true, nil
	}
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, fmt.Errorf("Error: negative exponents not supported.")
	} else if exponent == 0 {
		return 1, nil
	} else {
		result := 1
		for i := 0; i < exponent; i++ {
			result *= base
		}
		return result, nil
	}
}

func MakeCounter(start int) func() int {
	counter := start
	return func() int {
		counter += 1
		return counter
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func MakeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	accumulator := initial
	add = func(x int) {
		accumulator += x
	}
	subtract = func(x int) {
		accumulator -= x
	}
	get = func() int {
		return accumulator
	}
	return add, subtract, get
}

func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))

	for i, value := range nums {
		result[i] = operation(value)
	}

	return result
}

func Filter(nums []int, predicate func(int) bool) []int {
	result := []int{}

	for _, value := range nums {
		if predicate(value) {
			result = append(result, value)
		}
	}

	return result
}

func Reduce(nums []int, initial int, operation func(accumulator, current int) int) int {
	result := initial

	for _, value := range nums {
		result = operation(result, value)
	}

	return result
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func ExploreProcess() {
	// A process ID (PID) is a unique number the operating system gives
	// to a running program so it can be managed and identified.
	pid := os.Getpid()
	ppid := os.Getppid()

	fmt.Println("=== Process Information ===")
	fmt.Println("Current Process ID:", pid)
	fmt.Println("Parent Process ID:", ppid)

	// Create a slice of integers
	data := []int{1, 2, 3, 4, 5}

	// The address of the slice variable itself (slice header)
	fmt.Printf("Memory address of slice: %p\n", &data)

	// The address of the first element in the slice
	fmt.Printf("Memory address of first element: %p\n", &data[0])

	// Process isolation means each process has its own memory space.
	// Other processes cannot directly access this program's memory.
	fmt.Println("Note: Other processes cannot access these memory addresses due to process isolation")
}

// I dont think this will modify the original value, becuase the only gets a copy of the value and use it inside the function
// not modifying the original value
func DoubleValue(x int) int {
	return x * 2
}

// yes this will modify the original value, because we are passing the memory address of the variable
func DoublePointer(x *int) {
	*x = *x * 2
}

// this function creates a variable on the stack and returns its value
func CreateOnStack() int {
	y := 42
	return y
}

// create on the heap and return its address
func CreateOnHeap() *int {
	x := 15
	return &x
}

// swap two values using a temporary variable
func SwapValues(a, b int) (int, int) {
	swap := a
	a = b
	b = swap
	return a, b
}

// swap two values using pointers
func SwapPointers(a, b *int) (int, int) {
	swap := *a
	*a = *b
	*b = swap
	return *a, *b
}

// expirement
func AnalyzeEscape() {
	CreateOnStack()
	CreateOnHeap()
}

func main() {
	//x escaped to the heap, that happened because it was pass by a pointer so the variable is saved on the heap
	//rather than being deleted by the garbage collector
	AnalyzeEscape()

	fmt.Println("=== Process Information ===")
	ExploreProcess()

	fmt.Println("=== Math Operations ===")
	//calculate and print factorials
	fmt.Printf("Factorial(0): ")
	factorialResult1, err := Factorial(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(factorialResult1)

	fmt.Printf("Factorial(5): ")
	factorialResult2, err := Factorial(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(factorialResult2)

	fmt.Printf("Factorial(10): ")
	factorialResult3, err := Factorial(10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(factorialResult3)

	//check prime numbers
	fmt.Printf("IsPrime(17)")
	primeNumResult1, err := IsPrime(17)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(primeNumResult1)

	fmt.Printf("IsPrime(20)")
	primeNumResult2, err := IsPrime(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(primeNumResult2)

	fmt.Printf("IsPrime(25)")
	primeNumResult3, err := IsPrime(25)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(primeNumResult3)

	//power of numbers
	fmt.Printf("Power(2^8)")
	powerResult1, err := Power(2, 8)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(powerResult1)

	fmt.Printf("Power(5^3)")
	powerResult2, err := Power(5, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(powerResult2)

	cont1 := MakeCounter(0)
	cont2 := MakeCounter(100)

	//call function to show their indepency
	fmt.Println("=== Closure Demonstation ===")
	fmt.Println("Counter1: ", cont1())
	fmt.Println("Counter1: ", cont1())
	fmt.Println("Counter2: ", cont2())

	//doubler and tripler
	doubler := MakeMultiplier(2)
	tripler := MakeMultiplier(3)

	testNum := 5
	fmt.Println("\nDoubler and Tripler Demo:")
	fmt.Printf("Number: %d\n", testNum)
	fmt.Printf("Doubled: %d\n", doubler(testNum))
	fmt.Printf("Tripled: %d\n", tripler(testNum))

	//create a slice
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Use Apply to square all numbers
	fmt.Println("\nApply Demo (square all numbers):")
	fmt.Printf("Original slice: %v\n", slice)
	squared := Apply(slice, func(x int) int { return x * x })
	fmt.Printf("Squared: %v\n", squared)

	// Use Filter to get only even numbers
	fmt.Println("\nFilter Demo (even numbers):")
	evens := Filter(slice, func(x int) bool { return x%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evens)

	// Use Reduce to sum all numbers
	fmt.Println("\nReduce Demo (sum all numbers):")
	sum := Reduce(slice, 0, func(acc, curr int) int { return acc + curr })
	fmt.Printf("Sum of all numbers: %d\n", sum)

	// Use Compose to create a function that doubles then adds 10
	fmt.Println("\nCompose Demo (double then add 10):")
	doubleAndAdd10 := Compose(
		func(x int) int { return x + 10 },
		func(x int) int { return x * 2 },
	)
	testCompose := 5
	fmt.Printf("Input: %d\n", testCompose)
	fmt.Printf("Double then add 10: %d\n", doubleAndAdd10(testCompose))

	//pointer demo
	fmt.Println("=== Pointer Demonstration ===")

	// SwapValues demonstration
	fmt.Println("\nSwapValues Demo (pass by value - doesn't modify originals):")
	x, y := 10, 20
	fmt.Printf("Before: x=%d, y=%d\n", x, y)
	returnedX, returnedY := SwapValues(x, y)
	fmt.Printf("Returned: x=%d, y=%d\n", returnedX, returnedY)
	fmt.Printf("After: x=%d, y=%d (unchanged)\n", x, y)

	// SwapPointers demonstration
	fmt.Println("\nSwapPointers Demo (pass by reference - modifies originals):")
	a, b := 10, 20
	fmt.Printf("Before: a=%d, b=%d\n", a, b)
	returnedA, returnedB := SwapPointers(&a, &b)
	fmt.Printf("Returned: a=%d, b=%d\n", returnedA, returnedB)
	fmt.Printf("After: a=%d, b=%d (changed!)\n", a, b)

	
}
