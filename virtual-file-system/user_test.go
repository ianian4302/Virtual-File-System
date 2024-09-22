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
				username: "a",
			},
			want: false,
		},
		{
			name: "Test case 2",
			args: args{
				username: "test123",
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
