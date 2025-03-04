package shared

import "testing"

func TestTextContainerCreate(t *testing.T) {
	type args struct {
		header     string
		body       string
		lineLength int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01. Header and Body (16 chars per line)",
			args: args{
				header:     "HEADER",
				body:       "0102030405060708090a0b0c0d0e0f10",
				lineLength: 16,
			},
			want: "HEADER\n0102030405060708\n090a0b0c0d0e0f10\n",
		},
		{
			name: "02. Header and Body (8 chars per line)",
			args: args{
				header:     "HEADER",
				body:       "0102030405060708090a0b0c0d0e0f10",
				lineLength: 8,
			},
			want: "HEADER\n01020304\n05060708\n090a0b0c\n0d0e0f10\n",
		},
		{
			name: "03. Header and Body (4 chars per line)",
			args: args{
				header:     "HEADER",
				body:       "0102030405060708090a0b0c0d0e0f10",
				lineLength: 4,
			},
			want: "HEADER\n0102\n0304\n0506\n0708\n090a\n0b0c\n0d0e\n0f10\n",
		},
		{
			name: "04. Header and Body (1 char per line)",
			args: args{
				header:     "HEADER",
				body:       "0123456789abcdef",
				lineLength: 1,
			},
			want: "HEADER\n0\n1\n2\n3\n4\n5\n6\n7\n8\n9\na\nb\nc\nd\ne\nf\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TextContainerCreate(tt.args.header, tt.args.body, tt.args.lineLength); got != tt.want {
				t.Errorf("TextContainerCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}
