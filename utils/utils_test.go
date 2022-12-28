package utils

import "testing"

func TestGetLastIndex(t *testing.T) {
	type args struct {
		splits []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"U1test1",
			args{
				[]string{"a", "b", "c"},
			},
			"c",
		},
		{
			"U1test2",
			args{
				[]string{"1", "2", "3"},
			},
			"3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLastIndex(tt.args.splits); got != tt.want {
				t.Errorf("GetLastIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"U2Test1",
			args{
				"1",
			},
			1,
		},
		{
			"U2Test2",
			args{
				"2",
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitToInt(tt.args.s); got != tt.want {
				t.Errorf("SplitToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorCheck(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorCheck(tt.args.err)
		})
	}
}
