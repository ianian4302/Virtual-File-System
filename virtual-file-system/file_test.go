package virtualfilesystem

import (
	"testing"
)

func TestIsValidateFileName(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				filename: "0file",
			},
			want: false,
		},
		{
			name: "Test case 2",
			args: args{
				filename: " ",
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				filename: "file",
			},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{
				filename: "FILE",
			},
			want: true,
		},
		{
			name: "Test case 5",
			args: args{
				filename: "file1",
			},
			want: true,
		},
		{
			name: "Test case 6",
			args: args{
				filename: "file012345678901234567890",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidateFileName(tt.args.filename); got != tt.want {
				t.Errorf("IsValidateFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
