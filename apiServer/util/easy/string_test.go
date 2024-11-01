package easy

import (
	"fmt"
	"testing"
)

func TestVersionCompare(t *testing.T) {
	clientVersion := "0.1.1.1"
	serviceVersion := "0.1.1.2"
	fmt.Printf("clientVersion:%s,serviceVersion:%s,result:%t\r\n", clientVersion, serviceVersion, VersionCompare(clientVersion, serviceVersion))
	clientVersion = "0.1.1.0"
	serviceVersion = "0.1.1.0"
	fmt.Printf("clientVersion:%s,serviceVersion:%s,result:%t\r\n", clientVersion, serviceVersion, VersionCompare(clientVersion, serviceVersion))
	clientVersion = "0.3.1.1"
	serviceVersion = "0.4.1.2"
	fmt.Printf("clientVersion:%s,serviceVersion:%s,result:%t\r\n", clientVersion, serviceVersion, VersionCompare(clientVersion, serviceVersion))
	clientVersion = "0.1.2.1"
	serviceVersion = "0.1.2.2"
	fmt.Printf("clientVersion:%s,serviceVersion:%s,result:%t\r\n", clientVersion, serviceVersion, VersionCompare(clientVersion, serviceVersion))
	clientVersion = "1.1.1.1"
	serviceVersion = "1.1.1.2"
	fmt.Printf("clientVersion:%s,serviceVersion:%s,result:%t\r\n", clientVersion, serviceVersion, VersionCompare(clientVersion, serviceVersion))
	clientVersion = "2.1.1.1"
	serviceVersion = "1.1.1.2"
	fmt.Printf("clientVersion:%s,serviceVersion:%s,result:%t\r\n", clientVersion, serviceVersion, VersionCompare(clientVersion, serviceVersion))
}

func TestNumberToW(t *testing.T) {
	type args struct {
		src   interface{}
		fixed []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "10000",
			args: args{
				src:   10000,
				fixed: nil,
			},
			want: "1.00W",
		},
		{
			name: "12356",
			args: args{
				src:   12356,
				fixed: nil,
			},
			want: "1.24W",
		},
		{
			name: "12478",
			args: args{
				src:   12478,
				fixed: []int32{1},
			},
			want: "1.2W",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberToW(tt.args.src, tt.args.fixed...); got != tt.want {
				t.Errorf("NumberToW() = %v, want %v", got, tt.want)
			}
		})
	}
}
