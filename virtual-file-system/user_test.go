package virtualfilesystem

import (
	"testing"
)

func TestIsValidateUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//test lens
		{
			name: "Test case 1",
			args: args{
				username: "0user",
			},
			want: false,
		},
		{
			name: "Test case 2",
			args: args{
				username: "",
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				username: "user",
			},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{
				username: "USER",
			},
			want: true,
		},
		{
			name: "Test case 5",
			args: args{
				username: "user1",
			},
			want: true,
		},
		//test chars
		{
			name: "Test case 6",
			args: args{
				username: "user@",
			},
			want: false,
		},
		{
			name: "Test case 7",
			args: args{
				username: "user!",
			},
			want: false,
		},
		{
			name: "Test case 8",
			args: args{
				username: "user1",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidateUsername(tt.args.username); got != tt.want {
				t.Errorf("IsValidateUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
