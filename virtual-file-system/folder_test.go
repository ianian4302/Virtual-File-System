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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidateFolderName(tt.args.foldername); got != tt.want {
				t.Errorf("IsValidateFolderName() = %v, want %v", got, tt.want)
			}
		})
	}
}
