package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{name: "Factorial of 0", input: 0, want: 1, wantErr: false},
		{name: "Factorial of 5", input: 5, want: 120, wantErr: false},
		{name: "Factorial of negative number", input: -3, want: 0, wantErr: true},
		{name: "Factorial of 10", input: 10, want: 3628800, wantErr: false},
		{name: "Factorial of 1", input: 1, want: 1, wantErr: false},
		{name: "Factorial of 2", input: 2, want: 2, wantErr: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Factorial(tc.input)

			if (err != nil) != tc.wantErr {
				t.Errorf("Factorial() error == %v, wantErr %v", err, tc.wantErr)
				return
			}

			if got != tc.want {
				t.Errorf("Factorial() == %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{name: "Is prime of 0", input: 0, want: false, wantErr: true},
		{name: "Is prime of 1", input: 1, want: false, wantErr: true},
		{name: "Is prime of 2", input: 2, want: true, wantErr: false},
		{name: "Is prime of 3", input: 3, want: true, wantErr: false},
		{name: "Is prime of 4", input: 4, want: false, wantErr: false},
		{name: "Is prime of 5", input: 5, want: true, wantErr: false},
		{name: "Is prime of -1", input: -1, want: false, wantErr: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := IsPrime(tc.input)

			if (err != nil) != tc.wantErr {
				t.Errorf("IsPrime() error == %v, wantErr %v", err, tc.wantErr)
				return
			}

			if got != tc.want {
				t.Errorf("IsPrime() == %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	testCases := []struct {
		name     string
		base     int
		exponent int
		want     int
		wantErr  bool
	}{
		{name: "Power of 2^3", base: 2, exponent: 3, want: 8, wantErr: false},
		{name: "Power of 5^0", base: 5, exponent: 0, want: 1, wantErr: false},
		{name: "Power of 3^4", base: 3, exponent: 4, want: 81, wantErr: false},
		{name: "Power of 2^-2", base: 2, exponent: -2, want: 0, wantErr: true},
		{name: "Power of 10^2", base: 10, exponent: 2, want: 100, wantErr: false},
		{name: "Power of 7^1", base: 7, exponent: 1, want: 7, wantErr: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Power(tc.base, tc.exponent)

			if (err != nil) != tc.wantErr {
				t.Errorf("Power() error == %v, wantErr %v", err, tc.wantErr)
				return
			}

			if got != tc.want {
				t.Errorf("Power() == %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMakeCounter(t *testing.T) {
	t.Run("counter increments properly", func(t *testing.T) {
		counter := MakeCounter(0)

		if got := counter(); got != 1 {
			t.Errorf("First call: got %d, want 1", got)
		}
		if got := counter(); got != 2 {
			t.Errorf("Second call: got %d, want 2", got)
		}
		if got := counter(); got != 3 {
			t.Errorf("Third call: got %d, want 3", got)
		}
	})

	t.Run("counters are independent", func(t *testing.T) {
		counter1 := MakeCounter(0)
		counter2 := MakeCounter(10)

		// Increment counter1 a few times
		counter1()
		counter1()

		// counter1 should be at 2, counter2 should be at 11
		if got := counter1(); got != 3 {
			t.Errorf("counter1 after 2 increments: got %d, want 3", got)
		}

		if got := counter2(); got != 11 {
			t.Errorf("counter2 first call from 10: got %d, want 11", got)
		}

		// Verify counter1 continues independently
		if got := counter1(); got != 4 {
			t.Errorf("counter1 next call: got %d, want 4", got)
		}
	})
}

func TestMakeMulitiplier(t *testing.T) {
	testCases := []struct {
		name   string
		factor int
		input  int
		want   int
	}{
		{name: "Multiply by 2", factor: 2, input: 3, want: 6},
		{name: "Multiply by 5", factor: 5, input: 4, want: 20},
		{name: "Multiply by 0", factor: 0, input: 10, want: 0},
		{name: "Multiply by -1", factor: -1, input: 8, want: -8},
		{name: "Multiply by 3", factor: 3, input: -2, want: -6},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			multiplier := MakeMultiplier(tc.factor)
			got := multiplier(tc.input)
			if got != tc.want {
				t.Errorf("MakeMultiplier(%d)(%d) == %d, want %d", tc.factor, tc.input, got, tc.want)
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	t.Run("add increases the accumulator", func(t *testing.T) {
		add, _, get := MakeAccumulator(0)

		add(5)
		if got := get(); got != 5 {
			t.Errorf("After adding 5 to 0: got %d, want 5", got)
		}

		add(10)
		if got := get(); got != 15 {
			t.Errorf("After adding 10 to 5: got %d, want 15", got)
		}
	})

	t.Run("subtract decreases the accumulator", func(t *testing.T) {
		_, subtract, get := MakeAccumulator(100)

		subtract(30)
		if got := get(); got != 70 {
			t.Errorf("After subtracting 30 from 100: got %d, want 70", got)
		}

		subtract(20)
		if got := get(); got != 50 {
			t.Errorf("After subtracting 20 from 70: got %d, want 50", got)
		}
	})

	t.Run("add and subtract interact correctly", func(t *testing.T) {
		add, subtract, get := MakeAccumulator(10)

		// Start at 10
		if got := get(); got != 10 {
			t.Errorf("Initial value: got %d, want 10", got)
		}

		// Add 5: 10 + 5 = 15
		add(5)
		if got := get(); got != 15 {
			t.Errorf("After adding 5: got %d, want 15", got)
		}

		// Subtract 8: 15 - 8 = 7
		subtract(8)
		if got := get(); got != 7 {
			t.Errorf("After subtracting 8: got %d, want 7", got)
		}

		// Add 13: 7 + 13 = 20
		add(13)
		if got := get(); got != 20 {
			t.Errorf("After adding 13: got %d, want 20", got)
		}

		// Subtract 5: 20 - 5 = 15
		subtract(5)
		if got := get(); got != 15 {
			t.Errorf("After subtracting 5: got %d, want 15", got)
		}
	})

	t.Run("negative operations work correctly", func(t *testing.T) {
		add, subtract, get := MakeAccumulator(50)

		// Add negative number (like subtracting)
		add(-10)
		if got := get(); got != 40 {
			t.Errorf("After adding -10 to 50: got %d, want 40", got)
		}

		// Subtract negative number (like adding)
		subtract(-5)
		if got := get(); got != 45 {
			t.Errorf("After subtracting -5 from 40: got %d, want 45", got)
		}
	})
}

func TestApply(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		op   func(int) int
		want []int
	}{
		{"square", []int{1, 2, 3}, func(x int) int { return x * x }, []int{1, 4, 9}},
		{"double", []int{1, 2, 3}, func(x int) int { return x * 2 }, []int{2, 4, 6}},
		{"negate", []int{1, -2, 3}, func(x int) int { return -x }, []int{-1, 2, -3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.nums, tt.op)
			if len(got) != len(tt.want) {
				t.Errorf("Apply() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Apply() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		pred func(int) bool
		want []int
	}{
		{"even numbers", []int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 }, []int{2, 4}},
		{"positive numbers", []int{-2, -1, 0, 3}, func(x int) bool { return x > 0 }, []int{3}},
		{"greater than 10", []int{5, 10, 15, 20}, func(x int) bool { return x > 10 }, []int{15, 20}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.nums, tt.pred)
			if len(got) != len(tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Filter() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		init int
		op   func(int, int) int
		want int
	}{
		{"sum", []int{1, 2, 3}, 0, func(a, b int) int { return a + b }, 6},
		{"product", []int{1, 2, 3, 4}, 1, func(a, b int) int { return a * b }, 24},
		{"max", []int{3, 1, 5, 2}, -1, func(a, b int) int {
			if b > a {
				return b
			}
			return a
		}, 5},
		{"min", []int{3, 1, 5, 2}, 100, func(a, b int) int {
			if b < a {
				return b
			}
			return a
		}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.nums, tt.init, tt.op)
			if got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	tests := []struct {
		name string
		f    func(int) int
		g    func(int) int
		in   int
		want int
	}{
		{"double after add", func(x int) int { return x * 2 }, func(x int) int { return x + 1 }, 3, 8},
		{"square after add", func(x int) int { return x * x }, func(x int) int { return x + 2 }, 2, 16},
		{"negate after double", func(x int) int { return -x }, func(x int) int { return x * 2 }, 4, -8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := Compose(tt.f, tt.g)
			got := fn(tt.in)
			if got != tt.want {
				t.Errorf("Compose() = %v, want %v", got, tt.want)
			}
		})
	}
}

