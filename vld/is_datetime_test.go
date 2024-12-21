package vld

import "testing"

func TestIsDateTime(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid datetime",
			args: args{value: "2022-01-01 00:00:00"},
			want: true,
		},
		{
			name: "valid datetime",
			args: args{value: "2022-01-01T00:00:00"},
			want: true,
		},
		{
			name: "valid datetime with timezone",
			args: args{value: "2022-01-01T00:00:00Z"},
			want: true,
		},
		{
			name: "valid datetime with timezone",
			args: args{value: "2022-01-01 00:00:00Z"},
			want: true,
		},
		{
			name: "invalid datetime (no second colon)",
			args: args{value: "2022-01-01 00:00"},
			want: false,
		},
		{
			name: "invalid datetime (no colons)",
			args: args{value: "2022-01-01"},
			want: false,
		},
		{
			name: "invalid datetime (no dashes)",
			args: args{value: "20220101"},
			want: false,
		},
		{
			name: "invalid datetime (length)",
			args: args{value: "2022-01-01 00:00:00:00"},
			want: false,
		},
		{
			name: "invalid datetime (length)",
			args: args{value: "2022-01-01T00:00:00ZZ"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDateTime(tt.args.value); got != tt.want {
				t.Errorf("IsDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
