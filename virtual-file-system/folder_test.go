package virtualfilesystem

import (
	"testing"
)

func TestIsValidateFolderName(t *testing.T) {
	type args struct {
		foldername string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				foldername: "0folder",
			},
			want: false,
		},
		{
			name: "Test case 2",
			args: args{
				foldername: "",
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				foldername: "folder",
			},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{
				foldername: "FOLDER",
			},
			want: true,
		},
		{
			name: "Test case 5",
			args: args{
				foldername: "folder1",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidateFolderName(tt.args.foldername); got != tt.want {
				t.Errorf("IsValidateFolderName() = %v, want %v", got, tt.want)
			}
		})
	}
}
