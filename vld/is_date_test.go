package vld

import "testing"

func TestIsDate(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid date",
			args: args{value: "2022-01-01"},
			want: true,
		},
		{
			name: "invalid date (length)",
			args: args{value: "2022-01-011"},
			want: false,
		},
		{
			name: "invalid date (length)",
			args: args{value: "2022-01-1"},
			want: false,
		},
		{
			name: "invalid date (length)",
			args: args{value: "2022-01-"},
			want: false,
		},
		{
			name: "invalid date (length)",
			args: args{value: "2022-01"},
			want: false,
		},
		{
			name: "invalid date (length)",
			args: args{value: "2022-1"},
			want: false,
		},
		{
			name: "invalid date (length)",
			args: args{value: "2022"},
			want: false,
		},
		// {
		// 	name: "invalid date",
		// 	args: args{value: "2022-13-01"},
		// 	want: false,
		// },
		// {
		// 	name: "invalid date",
		// 	args: args{value: "2022-01-32"},
		// 	want: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDate(tt.args.value); got != tt.want {
				t.Errorf("IsDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
