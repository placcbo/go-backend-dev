package calculator

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -4, -6, -10},
		{"mixed sign", -3, 7, 4},
		{"zeros", 0, 0, 0},
		{"large numbers", 1000000, 2000000, 3000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive result", 10, 3, 7},
		{"negative result", 3, 10, -7},
		{"subtract zero", 5, 0, 5},
		{"same numbers", 7, 7, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 4, 5, 20},
		{"multiply by zero", 9, 0, 0},
		{"both negative", -3, -4, 12},
		{"mixed sign", -3, 4, -12},
		{"multiply by one", 7, 1, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    float64
		wantErr bool
	}{
		{"normal", 10, 2, 5, false},
		{"divide by zero", 10, 0, 0, true},
		{"negative result", 10, -2, -5, false},
		{"fraction", 1, 4, 0.25, false},
		{"zero numerator", 0, 5, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Divide(%v, %v) error = %v; wantErr %v",
					tt.a, tt.b, err, tt.wantErr)
			}

			if !tt.wantErr {
				diff := got - tt.want
				if diff < -0.001 || diff > 0.001 {
					t.Errorf("Divide(%v, %v) = %v; want %v",
						tt.a, tt.b, got, tt.want)
				}
			}
		})
	}
}
