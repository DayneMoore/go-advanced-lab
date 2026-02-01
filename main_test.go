package main
import(
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{ name: "Factorial of 0", input: 0, want: 1, wantErr: false,},
		{ name: "Factorial of 5", input: 5, want: 120, wantErr: false,},
		{ name: "Factorial of negative number", input: -3, want: 0, wantErr: true,},
		{ name: "Factorial of 10", input: 10, want: 3628800, wantErr: false,},
		{ name: "Factorial of 1", input: 1, want: 1, wantErr: false,},
		{ name: "Factorial of 2", input: 2, want: 2, wantErr: false,},

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
		{ name: "Is prime of 0", input: 0, want: false, wantErr: true},
		{ name: "Is prime of 1", input: 1, want: false, wantErr: true},
		{ name: "Is prime of 2", input: 2, want: true, wantErr: false},
		{ name: "Is prime of 3", input: 3, want: true, wantErr: false},
		{ name: "Is prime of 4", input: 4, want: false, wantErr: false},
		{ name: "Is prime of 5", input: 5, want: true, wantErr: false},
		{ name: "Is prime of -1", input: -1, want: false, wantErr: true},
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
		{ name: "Power of 2^3", base: 2, exponent: 3, want: 8, wantErr: false,},
		{ name: "Power of 5^0", base: 5, exponent: 0, want: 1, wantErr: false,},
		{ name: "Power of 3^4", base: 3, exponent: 4, want: 81, wantErr: false,},
		{ name: "Power of 2^-2", base: 2, exponent: -2, want: 0, wantErr: true,},
		{ name: "Power of 10^2", base: 10, exponent: 2, want: 100, wantErr: false,},
		{ name: "Power of 7^1", base: 7, exponent: 1, want: 7, wantErr: false,},
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
