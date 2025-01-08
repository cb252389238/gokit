package easy

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	type args struct {
		number float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test",
			args: args{
				number: 1.0,
			},
			want: 1.0,
		},
		{
			name: "test",
			args: args{
				number: -1.0,
			},
			want: 1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.number); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCeil2(t *testing.T) {
	t.Log(Ceil(1.1))
	t.Log(Ceil(1.5))
	t.Log(Ceil(1.9))
	t.Log(Ceil(1.0))
}

func TestFloor(t *testing.T) {
	t.Log(Floor(1.1))
	t.Log(Floor(1.5))
	t.Log(Floor(1.9))
	t.Log(Floor(1.0))
	t.Log(Floor(-1.1))
}

func TestMax(t *testing.T) {
	t.Log(Max(1, 2))
	t.Log(Max(1, 2, 3))
	t.Log(Max(1.0, 2.0, 3.0))
	t.Log(Max(1.0, 2.0, 3.0, 4.0))
	t.Log(Max(1.0, 2.0, 3.0, 4.0, 5.0))
	t.Log(Max(1.0, 2.0, 3.0, 4.0, 5.0, 6.0))
	t.Log(Max(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0))
}

func TestMin(t *testing.T) {
	t.Log(Min(1, 2))
	t.Log(Min(1, 2, 3))
	t.Log(Min(1.0, 2.0, 3.0))
	t.Log(Min(1.0, 2.0, 3.0, 4.0))
	t.Log(Min(1.0, 2.0, 3.0, 4.0, 5.0))
	t.Log(Min(1.0, 2.0, 3.0, 4.0, 5.0, 6.0))
	t.Log(Min(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0))
}

func TestGenerateRandomPrime(t *testing.T) {
	prime, err := GenerateRandomPrime(32)
	if err == nil {
		fmt.Println(IsPrime(prime))
	}
}

func BenchmarkGenerateRandomPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateRandomPrime(32)
	}
}
