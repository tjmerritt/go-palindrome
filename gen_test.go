package palindrome

import "testing"

func Test_multiplier(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
                {
                        name: "basic test",
                        args: args{
                                n: 2,
                        },
                        want:  100,
                },
                {
                        name: "negative value",
                        args: args{
                                n: -1,
                        },
                        want:  -1,
                },
                {
                        name: "zero value",
                        args: args{
                                n: 0,
                        },
                        want:  1,
                },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplier(tt.args.n); got != tt.want {
				t.Errorf("multiplier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstruct(t *testing.T) {
	type args struct {
		l int
		m int
		x int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
                {
                        name: "basic test",
                        args: args{
                                l: 82356,
                                m: 4,
                                x: 0,
                                c: 5,
                        },
                        want: 82356465328,
                },
                {
                        name: "no middle digit",
                        args: args{
                                l: 82356,
                                m: -1,
                                x: 0,
                                c: 5,
                        },
                        want: 8235665328,
                },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Construct(tt.args.l, tt.args.m, tt.args.x, tt.args.c); got != tt.want {
				t.Errorf("Construct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrev(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
                {
                        name: "basic test",
                        args: args{
                                n: 1287237,
                        },
                        want: 1286821,
                },
                {
                        name: "middle digit underflow",
                        args: args{
                                n: 1280237,
                        },
                        want: 1279721,
                },
                {
                        name: "no middle digit",
                        args: args{
                                n: 128237,
                        },
                        want: 127721,
                },
                {
                        name: "left side underflow",
                        args: args{
                                n: 1000001,
                        },
                        want: 999999,
                },
                {
                        name: "left side underflow, no middle digit",
                        args: args{
                                n: 100001,
                        },
                        want: 99999,
                },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prev(tt.args.n); got != tt.want {
				t.Errorf("Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNext(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
                {
                        name: "basic test",
                        args: args{
                                n: 1287237,
                        },
                        want: 1287821,
                },
                {
                        name: "middle digit underflow",
                        args: args{
                                n: 1289837,
                        },
                        want: 1290921,
                },
                {
                        name: "no middle digit",
                        args: args{
                                n: 128837,
                        },
                        want: 129921,
                },
                {
                        name: "left side overflow",
                        args: args{
                                n: 99999,
                        },
                        want: 100001,
                },
                {
                        name: "left side overflow, no middle digit",
                        args: args{
                                n: 999999,
                        },
                        want: 1000001,
                },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Next(tt.args.n); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
