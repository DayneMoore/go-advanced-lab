package main
import(
	"fmt"
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

func main() {
	add, sub, get := MakeAccumulator(100)
	add(50)
	fmt.Println(get())
	sub(30)
	fmt.Println(get())
}