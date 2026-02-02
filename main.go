package main
import(
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
	}	else if n == 2 {
		return true, nil
		}	else if n%2 == 0 {
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
	}	else if exponent == 0 {
		return 1, nil
	}	else {
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

//I dont think this will modify the original value, becuase the only gets a copy of the value and use it inside the function
//not modifying the original value
func DoubleValue(x int) int {
	return x * 2
}

//yes this will modify the original value, because we are passing the memory address of the variable
func DoublePointer(x *int) {
	*x = *x * 2
}

//this function creates a variable on the stack and returns its value
func CreateOnStack() int {
	y := 42
	return y
}

//create on the heap and return its address
func CreateOnHeap() *int {
	x := 515
	return &x
}

//swap two values using a temporary variable
func SwapValues(a, b int) (int, int) {
	swap := a
	a = b
	b = swap
	return a, b
}

//swap two values using pointers
func SwapPointers(a, b *int) (int, int) {
	swap := *a
	*a = *b
	*b = swap
	return *a, *b
}

func main() {
	ExploreProcess() 
}