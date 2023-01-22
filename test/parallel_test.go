package test

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal_1", args{1, 2}, 3},
		{"normal_2", args{2, 3}, 5},
		{"normal_3", args{3, 4}, 7},
	}
	for _, tt := range tests {

		tt := tt

		if testing.Short() {
			t.Log("skip test")
		}

		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := Parallel(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}