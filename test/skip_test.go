package test

import "testing"

func TestSkip(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "skip", args: args{a: 1, b: 2}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if testing.Short() {
				t.Skip("skipping test in short mode.")
			}

			if got := Skip(tt.args.a, tt.args.b); got != tt.want {
				t.Fatalf("add() = %v, want %v", got, tt.want)
			}

		})
	}
}